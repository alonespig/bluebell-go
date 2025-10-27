package repository

import (
	"bluebell/global"
	"bluebell/model"
	"context"
	"errors"
	"fmt"

	"go.uber.org/zap"
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

// GetPostListByPage 获取帖子列表
func GetPostListByPage(page, pageSize int) ([]*model.Post, error) {
	var posts []*model.Post
	if err := global.DB.Limit(pageSize).Offset((page - 1) * pageSize).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

// LikePost 点赞帖子
func LikePost(postID, userID int) error {
	likeKey := fmt.Sprintf("post:%d:like", postID)
	dislikeKey := fmt.Sprintf("post:%d:dislike", postID)
	// scoreKey := fmt.Sprintf("post:%d:score", postID)

	pipe := global.Redis.TxPipeline()
	ctx := context.Background()
	pipe.SAdd(ctx, likeKey, userID)
	pipe.SRem(ctx, dislikeKey, userID)
	// pipe.ZIncrBy(ctx, scoreKey, 1, userID)
	_, err := pipe.Exec(ctx)
	if err != nil {
		zap.S().Info("点赞失败", zap.Error(err))
		return errors.New("点赞失败")
	}
	return nil
}

// DislikePost 点踩帖子
func DislikePost(postID, userID int) error {
	likeKey := fmt.Sprintf("post:%d:like", postID)
	dislikeKey := fmt.Sprintf("post:%d:dislike", postID)
	// scoreKey := fmt.Sprintf("post:%d:score", postID)

	pipe := global.Redis.TxPipeline()
	ctx := context.Background()
	pipe.SRem(ctx, likeKey, userID)
	pipe.SAdd(ctx, dislikeKey, userID)
	// pipe.ZIncrBy(ctx, scoreKey, 1, userID)
	_, err := pipe.Exec(ctx)
	if err != nil {
		zap.S().Info("点踩失败", zap.Error(err))
		return errors.New("点踩失败")
	}
	return nil
}

// CancelVote 取消投票
func CancelVote(postID, userID int) error {
	likeKey := fmt.Sprintf("post:%d:like", postID)
	dislikeKey := fmt.Sprintf("post:%d:dislike", postID)

	pipe := global.Redis.TxPipeline()
	ctx := context.Background()
	pipe.SRem(ctx, likeKey, userID)
	pipe.SRem(ctx, dislikeKey, userID)
	_, err := pipe.Exec(ctx)
	if err != nil {
		zap.S().Info("取消投票失败", zap.Error(err))
		return errors.New("取消投票失败")
	}
	return nil
}

// GetPostVote 获取点赞数量
func GetPostLikeCount(postID int) (int, error) {
	likeKey := fmt.Sprintf("post:%d:like", postID)

	likeCount, err := global.Redis.SCard(context.Background(), likeKey).Result()
	if err != nil {
		return 0, err
	}
	return int(likeCount), nil
}

// GetPostDislike 获取点踩数量
func GetPostDislikeCount(postID int) (int, error) {
	dislikeKey := fmt.Sprintf("post:%d:dislike", postID)

	dislikeCount, err := global.Redis.SCard(context.Background(), dislikeKey).Result()
	if err != nil {
		return 0, err
	}
	return int(dislikeCount), nil
}

// GetUserVote 获取用户对帖子的投票状态
func GetUserVote(userID, postID int) (int, error) {
	likeKey := fmt.Sprintf("post:%d:like", postID)
	dislikeKey := fmt.Sprintf("post:%d:dislike", postID)

	like, err := global.Redis.SIsMember(context.Background(), likeKey, userID).Result()
	if err != nil {
		return 0, err
	}
	if like {
		return 1, nil
	}
	dislike, err := global.Redis.SIsMember(context.Background(), dislikeKey, userID).Result()
	if err != nil {
		return 0, err
	}
	if dislike {
		return -1, nil
	}
	return 0, nil
}

// GetPostListByUserID 获取用户发布的帖子列表
func GetPostListByUserID(userID int) ([]*model.Post, error) {
	var posts []*model.Post
	if err := global.DB.Where("author_id = ?", userID).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

// UpdatePost 更新帖子
func UpdatePost(authorID, postID int, title, content string) error {
	var count int64
	if err := global.DB.Model(&model.Post{}).Where("id = ? AND author_id = ?", postID, authorID).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return errors.New("帖子不存在")
	}
	return global.DB.Model(&model.Post{}).Where("id = ? AND author_id = ?", postID, authorID).Updates(map[string]interface{}{
		"title":   title,
		"content": content,
	}).Error
}

// DeletePost 删除帖子
func DeletePost(authorID, postID int) error {
	var count int64
	if err := global.DB.Model(&model.Post{}).Where("id = ? AND author_id = ?", postID, authorID).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return errors.New("帖子不存在")
	}
	return global.DB.Model(&model.Post{}).Where("id = ? AND author_id = ?", postID, authorID).Delete(&model.Post{}).Error
}
