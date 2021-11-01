package auth

import (
	"time"
	"strconv"
	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userId int, username string, roles []string) (string, error) {
	var err error
	claims := jwt.MapClaims{}
	claims["userId"] = strconv.Itoa(userId)
	claims["username"] = username
	claims["roles"] = roles
	claims["expiresAt"] = time.Now().Add(time.Minute * 60).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte("SECRET_KEY")) // exceptionnally put in code for more simplicity
	if err != nil {
	   return "", err
	}
	return token, nil
}