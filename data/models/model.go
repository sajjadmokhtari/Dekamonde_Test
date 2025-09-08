package models

import "time"

type User struct {
    ID        uint   `gorm:"primaryKey;autoIncrement"`
    Phone     string `gorm:"unique;not null"`
    CreatedAt time.Time
}
