package froms

import "time"

type CreatePostForm struct {
	CommunityID int    `json:"community_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
}

type PostDetailResponse struct {
	AuthorName   string                  `json:"author_name"`
	LikeCount    int                     `json:"like_count"`
	DislikeCount int                     `json:"dislike_count"`
	VoteNum      int                     `json:"vote_num"`
	ID           int                     `json:"id"`
	AuthorID     int                     `json:"author_id"`
	CommunityID  int                     `json:"community_id"`
	Status       int                     `json:"status"`
	Title        string                  `json:"title"`
	Content      string                  `json:"content"`
	CreatedAt    time.Time               `json:"created_at"`
	Community    CommunityDetailResponse `json:"community"`
}

type VotePostForm struct {
	PostID    int `json:"post_id" binding:"required"`
	Direction int `json:"direction" binding:"oneof=1 0 -1"`
}

type PostInfoResponse struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	VoteNum int    `json:"vote_num"`
}

type UpdatePostForm struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type GetPostListStatusForm struct {
	PostIDs []int `json:"post_ids"`
}
