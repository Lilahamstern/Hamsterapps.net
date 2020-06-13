package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/lilahamstern/hamsterapps.net/server/internal/users"
	"log"
	"os"
	"time"
)

var SecretKey = []byte(os.Getenv("SECRET"))

func GenerateToken(user users.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, getClaims(user))

	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal("Error occurred on generation of Token")
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenStr string) (users.User, error) {
	token, err := jwt.Parse(tokenStr, keyFunc)

	user := users.User{}

	if err != nil {
		return user, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user.Email = claims["email"].(string)
		// TODO: parse role
		return user, nil
	}

	return user, err
}

func getClaims(user users.User) jwt.Claims {
	claims := jwt.MapClaims{}

	claims["email"] = user.Email
	// TODO: Add rols at a later point
	claims["role"] = "tester"
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	return claims
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	return SecretKey, nil
}
