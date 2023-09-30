package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "task-5-pbi-btpns-mdafir/database"
    "task-5-pbi-btpns-mdafir/app"
    "task-5-pbi-btpns-mdafir/models"
)

type PhotoController struct{}

// Create digunakan untuk menambahkan foto profil oleh pengguna yang telah login
func (p *PhotoController) Create(c *gin.Context) {
    currentUser := c.MustGet("user").(models.User)
    var photo app.PhotoAdded

    if err := c.ShouldBindJSON(&photo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Dapatkan koneksi database
    db, err := database.ConnectDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }

    // Konversi struct dari app.PhotoAdded ke models.Photo
    photoModel := models.Photo{
        Title:    photo.Title,
        Caption:  photo.Caption,
        PhotoUrl: photo.PhotoUrl,
        UserID:   currentUser.ID,
    }

    if err := db.Create(&photoModel).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }

    c.JSON(http.StatusCreated, photoModel)
}

// GetAll digunakan untuk mendapatkan semua foto
func (p *PhotoController) GetAll(c *gin.Context) {
    // Dapatkan koneksi database
    db, err := database.ConnectDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }

    var photos []models.Photo
    if err := db.Find(&photos).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }

    c.JSON(http.StatusOK, photos)
}

// Update digunakan untuk mengubah foto oleh pemiliknya
func (p *PhotoController) Update(c *gin.Context) {
    currentUser := c.MustGet("user").(models.User)
    photoID := c.Param("photoId")

    var updatedPhoto app.PhotoUpdate
    if err := c.ShouldBindJSON(&updatedPhoto); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Dapatkan koneksi database
    db, err := database.ConnectDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }

    var existingPhoto models.Photo
    if err := db.Where("id = ? AND user_id = ?", photoID, currentUser.ID).First(&existingPhoto).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found or unauthorized"})
        return
    }

    // Perbarui data foto dalam database
    existingPhoto.Title = updatedPhoto.Title
    existingPhoto.Caption = updatedPhoto.Caption
    existingPhoto.PhotoUrl = updatedPhoto.PhotoUrl

    db.Save(&existingPhoto)
    c.JSON(http.StatusOK, existingPhoto)
}

// Delete digunakan untuk menghapus foto oleh pemiliknya
func (p *PhotoController) Delete(c *gin.Context) {
    currentUser := c.MustGet("user").(models.User)
    photoID := c.Param("photoId")

    // Dapatkan koneksi database
    db, err := database.ConnectDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }

    var existingPhoto models.Photo
    if err := db.Where("id = ? AND user_id = ?", photoID, currentUser.ID).First(&existingPhoto).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found or unauthorized"})
        return
    }

    db.Delete(&existingPhoto)
    c.JSON(http.StatusNoContent, gin.H{})
}