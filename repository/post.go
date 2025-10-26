package repository

import (
	"bluebell/global"
	"bluebell/model"
)

// CreatePost 创建帖子
func CreatePost(post *model.Post) error {
	return global.DB.Create(post).Error
}

// GetPostDetail 获取帖子详情
func GetPostDetail(postID int) (*model.Post, error) {
	var post model.Post
	if err := global.DB.Where("id = ?", postID).First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}
