package database

import (
    "task-5-pbi-btpns-mdafir/models"
    "gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
    db.AutoMigrate(&models.User{})
    db.AutoMigrate(&models.Photo{})
}