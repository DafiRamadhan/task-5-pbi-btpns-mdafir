package helpers

import (
    "github.com/dgrijalva/jwt-go"
    "time"
)

var jwtKey = []byte("your_secret_key")

type CustomClaims struct {
    UserID string `json:"user_id"`
    jwt.StandardClaims
}

func GenerateToken(userID string) (string, error) {
    claims := &CustomClaims{
        UserID: userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

func VerifyToken(tokenString string) (*CustomClaims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
        return claims, nil
    }

    return nil, err
}