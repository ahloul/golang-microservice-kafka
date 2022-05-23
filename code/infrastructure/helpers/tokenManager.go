package helpers

import (
	"errors"
	"fmt"
	"github.com/ahloul/loyalty-reports/infrastructure/config"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	viper := config.NewViper()
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.App.JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

//ExtractToken get the token from the request body
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func ExtractUserID(token *jwt.Token) (string, error) {

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId, userOk := claims["sub"].(string)
		fmt.Println(claims)

		if !ok || !userOk {
			return "", errors.New("unauthorized")
		} else {
			return userId, nil
		}
	}
	return "", errors.New("something went wrong")
}

func ExtractUserIDFromToken(r *http.Request) (string, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return "", err
	}
	acc, err := ExtractUserID(token)
	if err != nil {
		return "", err
	}
	return acc, nil
}
