package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func CreateToken(userid string) (string, error) {

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userid"] = userid
	claims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("ACCESS_KEY")))
}

func VerifyTokenHeader(c echo.Context) string {

	tokenData := c.Request().Header.Get("xauth")
	token, err := jwt.Parse(tokenData, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_KEY")), nil
	})
	if err != nil {
		return "unknown"
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		userid := fmt.Sprintf("%v", claims["userid"])
		return userid
	}
	return "unknown"
}

func VerifyTokenInBody(xauth string) string {

	tokenData := xauth
	token, err := jwt.Parse(tokenData, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_KEY")), nil
	})
	if err != nil {
		fmt.Println(err)
		return "unknown"
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		oauthid := fmt.Sprintf("%v", claims["oauthid"])
		return oauthid
	}
	return "unknown"
}
