package model

import "time"

type Community struct {
	ID            int       `json:"id" gorm:"primaryKey;autoIncrement"`
	CommunityID   int       `json:"community_id" gorm:"not null;unique"`
	CommunityName string    `json:"community_name" gorm:"type:varchar(128);not null;unique"`
	Introduction  string    `json:"introduction" gorm:"type:varchar(256);not null;default:''"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
