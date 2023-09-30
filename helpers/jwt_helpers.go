package helpers

import (
    "github.com/dgrijalva/jwt-go"
    "time"
)

var jwtKey = []byte("your_secret_key") // Ganti dengan kunci rahasia yang sesuai

// CustomClaims berisi informasi yang akan disimpan dalam token JWT
type CustomClaims struct {
    UserID string `json:"user_id"`
    jwt.StandardClaims
}

// GenerateToken digunakan untuk membuat token JWT
func GenerateToken(userID string) (string, error) {
    claims := &CustomClaims{
        UserID: userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token berlaku selama 1 hari
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

// VerifyToken digunakan untuk memverifikasi token JWT
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