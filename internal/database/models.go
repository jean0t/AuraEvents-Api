package database

import (
    "time"
    "gorm.io/gorm"
)


type User struct {
    gorm.Model
    Name string `gorm:"not null"`
    Email string `gorm:"not null;unique"`
    Password string `gorm:"not null"`
    Role string `gorm:"not null"` // can only be admin or user
}


type Event struct {
    gorm.Model
    OwnerID uint
    Owner User `gorm:"foreignKey:OwnerID"`
    Capacity uint
    Date time.Time
}


type Registration struct {
    gorm.Model
    UserID uint `gorm:"not null"`
    User User `gorm:"foreignKey:UserID"`
    EventID uint `gorm:"not null"`
    Event Event `gorm:"foreignKey:EventID"`
    WillAttend bool `gorm:"not null"`
}
