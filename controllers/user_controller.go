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

// Register digunakan untuk membuat akun pengguna baru
func (u *UserController) Register(c *gin.Context) {
	var user app.UserRegister
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password sebelum menyimpannya ke database
	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Konversi struct dari app.UserRegister ke models.User
	userModel := models.User{
		Username: user.Username,
		Email:    user.Email,
		Password: hashedPassword,
	}

	// Dapatkan koneksi database
	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Simpan pengguna ke database
	if err := db.Create(&userModel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusCreated, userModel)
}

// Login digunakan untuk otentikasi pengguna
func (u *UserController) Login(c *gin.Context) {
	var user app.UserLogin
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Dapatkan koneksi database
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

	// Buat token JWT dan kirimkan sebagai respons
	userIDStr := strconv.FormatUint(uint64(userModel.ID), 10) // Konversi user.ID menjadi string
	token, err := helpers.GenerateToken(userIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Update digunakan untuk mengubah data pengguna
func (u *UserController) Update(c *gin.Context) {
	currentUser := c.MustGet("user").(models.User)
	var user app.UserUpdate
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Periksa apakah pengguna yang diotorisasi sedang mencoba memperbarui data pengguna lain
	if currentUser.Email != user.Email {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Dapatkan koneksi database
	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Perbarui data pengguna dalam database
	db.Model(&currentUser).Updates(models.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password, // Hash password menggunakan MD5 sebelum disimpan
	})

	c.JSON(http.StatusOK, currentUser)
}

// Delete digunakan untuk menghapus akun pengguna
func (u *UserController) Delete(c *gin.Context) {
	currentUser := c.MustGet("user").(models.User)

	// Dapatkan koneksi database
	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Hapus akun pengguna dari database
	db.Delete(&currentUser)
	c.JSON(http.StatusNoContent, gin.H{})
}