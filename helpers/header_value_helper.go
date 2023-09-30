package helpers

import (
    "github.com/gin-gonic/gin"
)

func GetHeaderValue(c *gin.Context, headerName string) string {
    value := c.GetHeader(headerName)
    return value
}