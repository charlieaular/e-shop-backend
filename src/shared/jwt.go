package shared

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type MyCustomClaims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

var jwtSecret string = "my_secret_key"

func GenerateToken(id int) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"] = id

	accessTokenExpireTime := time.Now().Add(time.Hour * 24).Unix()
	claims["exp"] = accessTokenExpireTime

	tokenString, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		panic(err.Error())
	}

	return tokenString
}

func ValidateToken(signedToken string) (*MyCustomClaims, error) {
	tokenWithoutBearer := strings.Split(signedToken, "Bearer ")[1]

	token, err := jwt.ParseWithClaims(tokenWithoutBearer, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}

}
