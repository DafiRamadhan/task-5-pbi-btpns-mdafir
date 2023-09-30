package models

import (
    "gorm.io/gorm"
)

type Photo struct {
    gorm.Model
    ID        uint   `gorm:"primaryKey" json:"id"`
    Title     string `json:"title"`
    Caption   string `json:"caption"`
    PhotoUrl  string `json:"photo_url"`
    UserID    uint   `json:"user_id"`
    User      User   `json:"user"`
}