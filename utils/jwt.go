package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
    UserID int `json:"user_id"`
    jwt.RegisteredClaims
}


var SecretKey []byte



func CreateToken(id int)(string,error){
	expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        UserID: id,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
        },
    }

	SecretKeyString := os.Getenv("SECRET_KEY")
	if SecretKeyString == "" {
		fmt.Println("Secret key not found")
	}
	
	SecretKey = []byte(SecretKeyString)

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString,err := token.SignedString(SecretKey)
	if err !=nil{
		return "",err
	}
	return tokenString,nil
}