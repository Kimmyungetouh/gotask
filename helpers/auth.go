package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(b))
}

func CreateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func ExtractToken(context *gin.Context) string {
	token := context.Query("token")

	if token != "" {
		return token
	}

	bearerToken := context.Request.Header.Get("Authorization")
	tokenSlice := strings.Split(bearerToken, " ")
	if len(tokenSlice) == 2 {
		return tokenSlice[1]
	}
	return ""
}

func TokenValid(context *gin.Context) error {
	tokenString := ExtractToken(context)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET_KEY")), nil
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		Pretty(claims)
	}
	return nil
}

func ExtractTokenID(context *gin.Context) (uint, error) {
	tokenString := ExtractToken(context)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("bad signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("API_SECRET_KEY")), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		userID, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["userId"]), 10, 32)

		if err != nil {
			return 0, err
		}
		return uint(userID), err
	}

	return 0, nil

}
