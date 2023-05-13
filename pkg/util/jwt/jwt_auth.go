package util

import (
	"errors"
	"io/ioutil"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type myClaims struct {
	Username string `json:"username"`
	UserId   int    `json:"id"`
	jwt.StandardClaims
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
func TokenAuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.Request.Header.Get("Authorization")

		if authHeader == "" {
			respondWithError(c, 401, "Authorization header is required")
			return
		}

		splitToken := strings.Split(authHeader, "Bearer ")
		publicKeyPath := "./jwtRS256.key.pub"
		token := splitToken[1]
		isValid, err := ValidateToken(token, publicKeyPath)
		if err != nil {
			log.Fatal(err)
		}
		if !isValid {
			respondWithError(c, 401, "Invalid API token")
		}
		c.Next()
	}
}
func ValidateToken(token, publicKeyPath string) (bool, error) {
	keyData, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return false, err
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		return false, err
	}
	parts := strings.Split(token, ".")
	err = jwt.SigningMethodRS256.Verify(strings.Join(parts[0:2], "."), parts[2], key)
	if err != nil {
		return false, nil
	}
	_, err = jwt.ParseWithClaims(token, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return false, nil
	}

	return true, nil
}
func GenerateToken(privateKeyPath string) (string, error) {
	privateKeyData, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return "", err
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEMWithPassword(privateKeyData, "empreinte")
	if err != nil {
		return "", err
	}
	newToken := jwt.NewWithClaims(jwt.SigningMethodRS256, myClaims{})

	// Set the signing key and sign the token
	newTokenString, err := newToken.SignedString(privateKey)
	if err != nil {
		return "", errors.New("jwt token generation failed")
	}

	return newTokenString, nil

}
