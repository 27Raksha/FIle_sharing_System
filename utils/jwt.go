package utils

import (
    "time"
    "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("your_jwt_secret")

type Claims struct {
    Email string `json:"email"`
    jwt.StandardClaims
}


func GenerateJWT(email string) (string, error) {
    expirationTime := time.Now().Add(72 * time.Hour)
    claims := &Claims{
        Email: email,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}


func ValidateJWT(tokenStr string) (*Claims, bool) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })
    if err != nil {
        return nil, false
    }
    return claims, token.Valid
}

