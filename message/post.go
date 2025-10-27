package message

import (
	"bluebell/global"
	"context"
	"encoding/json"

	"github.com/apache/rocketmq-client-go/v2/primitive"
)

// LikeMessage 点赞消息
type LikeMessage struct {
	PostID int `json:"post_id"`
	UserID int `json:"user_id"`
	Action int `json:"action"` // 1=赞, -1=踩, 0=取消
}

func SendLikeEvent(ctx context.Context, postID, userID, action int) error {
	data, err := json.Marshal(LikeMessage{
		PostID: postID,
		UserID: userID,
		Action: action,
	})
	if err != nil {
		return err
	}
	_, err = global.Producer.SendSync(ctx, primitive.NewMessage("vote_post", data))
	return err
}
