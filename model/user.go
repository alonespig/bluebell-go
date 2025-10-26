package model

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    int       `json:"user_id" gorm:"not null;unique"`
	Username  string    `json:"username" gorm:"type:varchar(64);not null;unique"`
	Password  string    `json:"password" gorm:"type:varchar(64);not null"`
	Email     string    `json:"email" gorm:"type:varchar(64);"`
	Gender    int       `json:"gender" gorm:"type:int;not null;default:0"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
