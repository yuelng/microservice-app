package auth

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

// use standardClaims, only save id
func GenerateJWT(id, duration int64, secret string) string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:        strconv.FormatInt(id, 10),
		ExpiresAt: time.Now().AddDate(0, 0, int(duration)).Unix(),
	})

	// time.Now().Add(time.Hour * 24).Unix()
	// token := jwt.New(jwt.SigningMethodHS256)
	// token.Claims["admin"] = true

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(fmt.Errorf("generateJWT failed: %v", err))
	}

	return tokenString
}

func ParseJWT(tokenString, secret string) (int64, error) {
	var claims jwt.StandardClaims
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return claims.Id, err
	}

	return claims.Id, nil
}
