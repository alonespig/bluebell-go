package model

import "time"

type Post struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	PostID      int       `json:"post_id" gorm:"unique"`
	Title       string    `json:"title" gorm:"type:varchar(128);not null;index"`
	Content     string    `json:"content" gorm:"type:text;not null"`
	AuthorID    int       `json:"author_id" gorm:"not null"`
	CommunityID int       `json:"community_id" gorm:"not null;index"`
	Status      int       `json:"status" gorm:"type:int;not null;default:1"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
