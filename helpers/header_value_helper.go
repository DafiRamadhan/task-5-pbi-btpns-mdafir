package helpers

import (
    "github.com/gin-gonic/gin"
)

// GetHeaderValue digunakan untuk mendapatkan nilai header tertentu dari request
func GetHeaderValue(c *gin.Context, headerName string) string {
    value := c.GetHeader(headerName)
    return value
}