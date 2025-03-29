package models

import "time"

type User struct {
    ID         uint      `gorm:"primaryKey"`
    Username   string
    Email      string    `gorm:"unique"`
    Password   string
    DateJoined time.Time
    Location   string
}

type AstronomicalEvent struct {
    ID          uint      `gorm:"primaryKey"`
    Name        string
    Type        string
    Description string
    DateTime    time.Time
    Visibility  string
}

type Observation struct {
    ID          uint      `gorm:"primaryKey"`
    UserID      uint
    EventID     uint
    Date        time.Time
    Location    string
    Description string
}

type Comment struct {
    ID             uint      `gorm:"primaryKey"`
    ObservationID  uint
    UserID         uint
    Content        string
    CreatedAt      time.Time
}

type Notification struct {
    ID               uint      `gorm:"primaryKey"`
    UserID           uint
    EventID          uint
    NotificationDate time.Time
}
