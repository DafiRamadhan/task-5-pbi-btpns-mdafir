package models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    ID       uint   `gorm:"primaryKey" json:"id"`
    Username string `json:"username"`
    Email    string `gorm:"unique" json:"email"`
    Password string `json:"-"`
    Photos   []Photo
}