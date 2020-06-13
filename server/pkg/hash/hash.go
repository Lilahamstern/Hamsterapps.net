package hash

import (
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
)

var Cost = os.Getenv("HASH_COST")

func HashPassword(password string) (string, error) {
	cost, _ := strconv.ParseInt(Cost, 10, 32)

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), int(cost))
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
