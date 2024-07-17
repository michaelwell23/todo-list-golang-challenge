package utils

import (
    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
    "log"
    "mytodoapp/src/models"
)

var DB *gorm.DB

func InitDB() {
    var err error
    DB, err = gorm.Open("sqlite3", "test.db")
    if err != nil {
        log.Fatal("failed to connect database", err)
    }

    DB.AutoMigrate(&models.Task{})
}
