package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"task-5-pbi-btpns-mdafir/database"
	"task-5-pbi-btpns-mdafir/helpers"
	"task-5-pbi-btpns-mdafir/models"
	"task-5-pbi-btpns-mdafir/app"
)

type UserController struct{}

func (u *UserController) Register(c *gin.Context) {
	var user app.UserRegister
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	userModel := models.User{
		Username: user.Username,
		Email:    user.Email,
		Password: hashedPassword,
	}

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if err := db.Create(&userModel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusCreated, userModel)
}

func (u *UserController) Login(c *gin.Context) {
	var user app.UserLogin
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	var userModel models.User
	if err := db.Where("email = ?", user.Email).First(&userModel).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if !helpers.CheckPasswordHash(user.Password, userModel.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	userIDStr := strconv.FormatUint(uint64(userModel.ID), 10)
	token, err := helpers.GenerateToken(userIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (u *UserController) Update(c *gin.Context) {
	currentUser := c.MustGet("user").(models.User)
	var user app.UserUpdate
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if currentUser.Email != user.Email {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	db.Model(&currentUser).Updates(models.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	})

	c.JSON(http.StatusOK, currentUser)
}

func (u *UserController) Delete(c *gin.Context) {
	currentUser := c.MustGet("user").(models.User)

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	db.Delete(&currentUser)
	c.JSON(http.StatusNoContent, gin.H{})
}