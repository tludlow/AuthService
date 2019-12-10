package utils

import "golang.org/x/crypto/bcrypt"

//bcrpyrCost - The number of times to run the algorithm on the password.
//2^11 times - https://security.stackexchange.com/questions/17207/recommended-of-rounds-for-bcrypt
const bcryptCost = 11

//HashPassword - Take a plaintext user password and generate a hashed version of this password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	return string(bytes), err
}

//CheckPasswordHash - Compare a plaintext password against the hashed version of the password.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
