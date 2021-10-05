package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userid string) (string, error) {

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userid"] = userid
	claims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("11"))
}

func VerifyToken(xauth string) string {

	tokenData := xauth
	token, err := jwt.Parse(tokenData, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("11"), nil
	})
	if err != nil {
		fmt.Println(err)
		return "unknown"
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		userid := fmt.Sprintf("%v", claims["userid"])
		return userid
	}
	return "unknown"
}