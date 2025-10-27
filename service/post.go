package service

import (
	"bluebell/froms"
	"bluebell/message"
	"bluebell/model"
	"bluebell/repository"
	"context"
	"errors"
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

	likeCount, err := repository.GetPostLikeCount(post.ID)
	if err != nil {
		return nil, err
	}
	dislikeCount, err := repository.GetPostDislikeCount(post.ID)
	if err != nil {
		return nil, err
	}
	return &froms.PostDetailResponse{
		AuthorName:   user.Username,
		LikeCount:    likeCount,
		DislikeCount: dislikeCount,
		VoteNum:      likeCount,
		ID:           post.ID,
		AuthorID:     post.AuthorID,
		CommunityID:  post.CommunityID,
		Status:       post.Status,
		Title:        post.Title,
		Content:      post.Content,
		CreatedAt:    post.CreatedAt,
		Community: froms.CommunityDetailResponse{
			ID:           community.ID,
			Name:         community.CommunityName,
			Introduction: community.Introduction,
			CreatedAt:    community.CreatedAt,
		},
	}, nil
}

// GetPostListByPage 获取帖子列表
func GetPostListByPage(page, pageSize int) ([]*froms.PostDetailResponse, error) {
	posts, err := repository.GetPostListByPage(page, pageSize)
	if err != nil {
		return nil, err
	}
	var postDetailRespList []*froms.PostDetailResponse
	for _, post := range posts {
		community, err := repository.GetCommunityDetail(post.CommunityID)
		if err != nil {
			return nil, err
		}
		user, err := repository.GetUserByID(post.AuthorID)
		if err != nil {
			return nil, err
		}
		likeCount, err := repository.GetPostLikeCount(post.ID)
		if err != nil {
			return nil, err
		}
		dislikeCount, err := repository.GetPostDislikeCount(post.ID)
		if err != nil {
			return nil, err
		}
		postDetailRespList = append(postDetailRespList, &froms.PostDetailResponse{
			AuthorName:   user.Username,
			LikeCount:    likeCount,
			DislikeCount: dislikeCount,
			VoteNum:      likeCount,
			ID:           post.ID,
			AuthorID:     post.AuthorID,
			CommunityID:  post.CommunityID,
			Status:       post.Status,
			Title:        post.Title,
			Content:      post.Content,
			CreatedAt:    post.CreatedAt,
			Community: froms.CommunityDetailResponse{
				ID:           community.ID,
				Name:         community.CommunityName,
				Introduction: community.Introduction,
				CreatedAt:    community.CreatedAt,
			},
		})
	}
	return postDetailRespList, nil
}

// VoteForPost 投票
func VoteForPost(userID int, form froms.VotePostForm) error {
	err := message.SendLikeEvent(context.Background(), form.PostID, userID, form.Direction)
	if err != nil {
		return err
	}
	switch form.Direction {
	case 1:
		return repository.LikePost(form.PostID, userID)
	case -1:
		return repository.DislikePost(form.PostID, userID)
	case 0:
		return repository.CancelVote(form.PostID, userID)
	default:
		return errors.New("direction is not valid")
	}
}

// GetUserVote 获取用户对帖子的投票状态
func GetUserVote(userID, postID int) (int, error) {
	return repository.GetUserVote(userID, postID)
}

// GetPostListByUserID 获取用户发布的帖子列表
func GetPostListByUserID(userID int) ([]*froms.PostInfoResponse, error) {
	posts, err := repository.GetPostListByUserID(userID)
	if err != nil {
		return nil, err
	}
	var postInfoList []*froms.PostInfoResponse
	for _, post := range posts {
		likeCount, err := repository.GetPostLikeCount(post.ID)
		if err != nil {
			return nil, err
		}
		postInfoList = append(postInfoList, &froms.PostInfoResponse{
			ID:      post.ID,
			Title:   post.Title,
			Content: post.Content,
			VoteNum: likeCount,
		})
	}
	return postInfoList, nil
}

// UpdatePost 更新帖子
func UpdatePost(authorID, postID int, title, content string) error {
	return repository.UpdatePost(authorID, postID, title, content)
}

// DeletePost 删除帖子
func DeletePost(authorID, postID int) error {
	return repository.DeletePost(authorID, postID)
}

// GetPostListByStatus 获取帖子列表状态
func GetPostListByStatus(postIDs []int, userID int) ([]int, error) {
	postsStatus := make([]int, len(postIDs))
	for i, postID := range postIDs {
		postsStatus[i], _ = repository.GetUserVote(userID, postID)
	}
	return postsStatus, nil
}
