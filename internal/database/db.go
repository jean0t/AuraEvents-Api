package database

import (
    "gorm.io/gorm"
)


type dbConnection interface {
    Open(cdn string) gorm.Dialector
}


func ConnectDB(connection dbConnection, cdn string) (*gorm.DB, error) {
    var (
        err error
        db *gorm.DB
    )

    db, err = gorm.Open(connection.Open(cdn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    return db, nil
}
