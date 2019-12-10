package communications

import (
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/tludlow/authservice/database"
	"github.com/tludlow/authservice/utils"
)

//SetupRoutes - Sets the specific routes and their HTTP verb to the specific handlers.
func SetupRoutes(router *gin.Engine) {
	router.POST("/user/create", CreateUser)
}

//CreateUser - Creates a new user if not existing, errors if existing or breaks user creation rules.
//POST - /user/create
func CreateUser(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirmPassword")

	if len(username) == 0 || len(email) == 0 || len(password) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username, email or password had length of 0",
		})
		return
	}

	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(email) > 254 || !emailRegex.MatchString(email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email structure invalid",
		})
		return
	}

	//TODO - May want to convert this to a []byte secure comparison.
	if password != confirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Passwords do not match",
		})
		return
	}

	//Find the number of users that have either the username provided or the email provided.
	userCount, err := database.UsersByNameOrEmail(username, email)

	if err != nil {
		log.Fatal("Error checking for prexisting username/email combination. - " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error checking for username/email existence.",
		})
		return
	}

	if userCount != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username or email already taken.",
		})
		return
	}

	//Encrypt the password using bcrypt.
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		log.Fatal(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Can't generate your user details.",
		})
		return
	}

	//Create the new user in the db.
	err = database.InsertNewUser(username, email, hashedPassword)
	if err != nil {
		log.Fatal(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Can't insert your details into the database",
		})
		return
	}

	//Return the new user info
	user, err := database.GetUserByUsername(username)
	if err != nil {
		log.Fatal(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Can't get new user details.",
		})
		return
	}

	//Everything has worked, return the new user.
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
