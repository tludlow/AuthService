package database

import (
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tludlow/authservice/models"
)

//UsersByNameOrEmail - Finds the amount of users existing with the specified username or email.
func UsersByNameOrEmail(searchUsername, searchEmail string) (int, error) {
	var usersCount int

	err := Database.QueryRow("SELECT COUNT(*) AS users_found FROM users WHERE username=? OR email=?", searchUsername, searchEmail).Scan(&usersCount)
	if err != nil {
		return -1, err
	}

	return usersCount, nil
}

//GetUserByUsername - Returns a user model from the username
func GetUserByUsername(username string) (models.User, error) {
	rows, err := Database.Query("SELECT username, email, created_at FROM users WHERE username=?", username)
	if err != nil {
		log.Fatal(err.Error())
		return models.User{}, errors.New("Error finding user")
	}
	defer rows.Close()

	//We have a result
	if rows.Next() {
		returnUser := models.User{}

		err = rows.Scan(&returnUser.Username, &returnUser.Email, &returnUser.CreatedAt)
		if err != nil {
			log.Fatal(err.Error())
			return models.User{}, errors.New("Can't get user information")
		}
		return returnUser, nil
	}

	//No result found
	return models.User{}, errors.New("No user information found from search username")
}

//InsertNewUser - Creates a new user in the database.
//*** Ensure that the password provided is encrypted using the util methods.
//Should return the GetUserByUsername() result to the user upon creation.
func InsertNewUser(username, email, password string) error {
	//Check that the user doesn't already exist.
	usersFound, err := UsersByNameOrEmail(username, email)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	if usersFound > 0 {
		return errors.New("User with that username or password already exists")
	}

	//Now create the user.
	stmt, err := Database.Prepare("INSERT INTO users(username, email, password) VALUES(?, ?, ?)")
	defer stmt.Close()
	if err != nil {
		log.Fatal(err.Error())
		return errors.New("Can't execute user insert query")
	}

	_, err = stmt.Exec(username, email, password)
	if err != nil {
		log.Fatal(err.Error())
		return errors.New("Cant insert user into database")
	}

	return nil
}
