package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Ensuring that hashing a password works as expected to generate secure password hashes.
func TestPasswordEncrypt(t *testing.T) {
	plaintextPassword := "random_passwordcool!"

	//Hash the password.
	hashedPassword, err := HashPassword(plaintextPassword)

	assert.Nil(t, err, "Error whilst hashing password")
	assert.NotNil(t, hashedPassword, "The hashed password is nil, it shouldnt be.")
	assert.NotEqual(t, plaintextPassword, hashedPassword, "Plaintext password equal to the hashed password")
	assert.GreaterOrEqual(t, len(hashedPassword), 30, "The hashed password isn't very long in length.")
}

//Ensuring that the comparison between a plaintext password and a hashed version of that password works.
func TestPasswordComparison(t *testing.T) {
	plaintextPassword := "another_raNDom_plaintexTpass123!"
	hashedPassword, err := HashPassword(plaintextPassword)

	assert.Nil(t, err, "Error when hashing password")
	assert.NotNil(t, hashedPassword, "The hashed password is nil, it shouldnt be.")

	//Compare the password with the known password and unrelated passwords.
	comparison := CheckPasswordHash(plaintextPassword, hashedPassword)
	assert.True(t, comparison, "The plaintext password is not the same as it's hashed equivalent")

	comparison2 := CheckPasswordHash("definetleynotthesame", hashedPassword)
	assert.False(t, comparison2, "The hash of another password matches that of other, not the same plaintext.")

	comparison3 := CheckPasswordHash("nother_raNDom_plaintexTpass123", hashedPassword)
	assert.False(t, comparison3, "The hash of another password matches that of other, not the same plaintext.")

	comparison4 := CheckPasswordHash("Another_raNDom_plaintexTpass123", hashedPassword)
	assert.False(t, comparison4, "The hash of another password matches that of other, not the same plaintext.")
}

//Ensuring that the password hashing cost is greater than, or equal to 10. The recommended amount for secure passwords.
func TestBcryptSecure(t *testing.T) {
	assert.GreaterOrEqual(t, bcryptCost, 10, "The bcrypt cost is less than 10. This is not recommended for secure hashing.")
}
