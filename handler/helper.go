package handler

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/samul333/go-auth/database"
	"golang.org/x/crypto/bcrypt"
)

func HashPasswords(password string) (string, error) {
	passwordBytes := []byte(password)

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(hashedPasswordBytes), nil

}

func GenerateJwtToken(user *database.User, jwtSecret string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	s, err := token.SignedString([]byte("secret"))

	if err != nil {
		return "", err
	}

	return s, nil

}

func DecodePassword(hashedPassword, currPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(currPassword))
	return err == nil
}
