package froms

import "time"

type CreatePostForm struct {
	CommunityID int    `json:"community_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
}

type PostDetailResponse struct {
	AuthorName  string                  `json:"author_name"`
	VoteNum     int                     `json:"vote_num"`
	ID          int                     `json:"id"`
	AuthorID    int                     `json:"author_id"`
	CommunityID int                     `json:"community_id"`
	Status      int                     `json:"status"`
	Title       string                  `json:"title"`
	Content     string                  `json:"content"`
	CreatedAt   time.Time               `json:"created_at"`
	Community   CommunityDetailResponse `json:"community"`
}
