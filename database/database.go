package database

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
    dsn := "root:@tcp(localhost:3306)/task-5-pbi-btpns?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    return db, nil
}