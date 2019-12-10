package utils

//Expirations expressed in hours, therefore 30 mins = 0.5
const accessTokenExpiration = 0.5
const refreshTokenExpiration = 24 * 7

//GenerateNewTokens - Generates a new set of tokens for a user. A refresh token and an access token.
func GenerateNewTokens() (string, string, error) {
	return "", "", nil
}

//GenerateAccessToken - Generate a new access token based on the provided refresh token.
func GenerateAccessToken(refreshToken string) (string, error) {
	return "", nil
}
