package util

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

var jwtKey []byte

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))
}

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(username, role string) (string, error) {
	claims := &Claims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)

}
