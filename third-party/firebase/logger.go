package firebase

import (
	"gorm.io/gorm"
)

var FirebaseLogger Logger

type Notification struct {
	gorm.Model
	//! YOUR NOTIFICATION MODEL
}

type UserNotification struct {
	gorm.Model
	//! YOUR USER NOTIFICATION MODEL
}

type Logger struct {
	db *gorm.DB
}

func NewLogger(db *gorm.DB) *Logger {
	return &Logger{db}
}
