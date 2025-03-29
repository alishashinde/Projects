package config

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "os"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := os.Getenv("DATABASE_URL")
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }

    DB = database
    log.Println("Database connection established")
}
