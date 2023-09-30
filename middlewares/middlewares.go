package middlewares

import (
    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
    "task-5-pbi-btpns-mdafir/helpers"
    "task-5-pbi-btpns-mdafir/models"
    "net/http"
    "strings"
    "strconv" // Import paket strconv untuk konversi tipe data
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        // Parse token
        tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
        token, err := jwt.ParseWithClaims(tokenString, &helpers.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
            return []byte("your_secret_key"), nil
        })

        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        if claims, ok := token.Claims.(*helpers.CustomClaims); ok && token.Valid {
            // Konversi claims.UserID dari string ke uint
            userID, err := strconv.ParseUint(claims.UserID, 10, 32)
            if err != nil {
                c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
                c.Abort()
                return
            }

            user := models.User{ID: uint(userID)}
            c.Set("user", user)
            c.Next()
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }
    }
}