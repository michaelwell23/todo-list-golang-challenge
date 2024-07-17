package models

import (
    "github.com/jinzhu/gorm"
)

type Task struct {
    gorm.Model
    Description string `json:"description"`
    Responsible string `json:"responsible"`
    Email       string `json:"email"`
    Completed   bool   `json:"completed"`
    RevertCount int    `json:"revert_count"`
}
