package infrastructure

import (
	"os"
	"time"
	"log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var jwtKey = []byte(func() string{
	err := godotenv.Load("config/.env")
	if err != nil{
		log.Fatalln("Error loading .env file: ", err)
	}
	s := os.Getenv("JWT_SECRET")
	if s == ""{
		panic("JWT_SECRET not set")
	}
	return s
}())

type Claims struct{
	Username string `json:"username"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(username, role string) (string,error){
	claims := Claims{
		Username: username,
		Role: role,
        RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString(jwtKey)
}

func ParseToken(tokenStr string) (*Claims, error){
	tok, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error){
			return jwtKey, nil
		})
	if err != nil{
		return nil, err
	}
	return tok.Claims.(*Claims), nil
}