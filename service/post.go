package service

import (
	"bluebell/froms"
	"bluebell/model"
	"bluebell/repository"
)

// CreatePost 创建帖子
func CreatePost(post *model.Post) error {
	return repository.CreatePost(post)
}

// GetPostDetail 获取帖子详情
func GetPostDetail(postID int) (*froms.PostDetailResponse, error) {
	post, err := repository.GetPostDetail(postID)
	if err != nil {
		return nil, err
	}
	community, err := repository.GetCommunityDetail(post.CommunityID)
	if err != nil {
		return nil, err
	}

	user, err := repository.GetUserByID(post.AuthorID)
	if err != nil {
		return nil, err
	}

	return &froms.PostDetailResponse{
		AuthorName:  user.Username,
		VoteNum:     0,
		ID:          post.ID,
		AuthorID:    post.AuthorID,
		CommunityID: post.CommunityID,
		Status:      post.Status,
		Title:       post.Title,
		Content:     post.Content,
		CreatedAt:   post.CreatedAt,
		Community: froms.CommunityDetailResponse{
			ID:           community.ID,
			Name:         community.CommunityName,
			Introduction: community.Introduction,
			CreatedAt:    community.CreatedAt,
		},
	}, nil
}
