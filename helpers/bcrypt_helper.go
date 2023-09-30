package helpers

import (
    "golang.org/x/crypto/bcrypt"
)

// HashPassword digunakan untuk meng-hash password pengguna
func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

// CheckPasswordHash digunakan untuk memeriksa apakah password cocok dengan hash yang disimpan
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}