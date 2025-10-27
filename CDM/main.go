package main

import (
	"bluebell/CDM/mysql"
	"bluebell/global"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func init() {
	global.InitConfig()
	global.InitLogger()
	mysql.Init()
	global.InitConsumer([]string{"172.27.49.67:9876"}, "vote_post", "vote_post")
}

// LikeMessage 点赞消息
type LikeMessage struct {
	PostID int `json:"post_id"`
	UserID int `json:"user_id"`
	Action int `json:"action"` // 1=赞, -1=踩, 0=取消
}

func handleLikeEvent(msg LikeMessage) error {
	sqlStr1 := `select count(*) from post_votes where post_id = ? and user_id = ?`
	var count int
	err := mysql.DB.Get(&count, sqlStr1, msg.PostID, msg.UserID)
	if err != nil {
		return err
	}

	if count > 0 {
		sqlStr2 := `update post_votes set vote = ? where post_id = ? and user_id = ?`
		_, err = mysql.DB.Exec(sqlStr2, msg.Action, msg.PostID, msg.UserID)
		if err != nil {
			return err
		}
	} else {
		sqlStr3 := `insert into post_votes (post_id, user_id, vote) values (?, ?, ?)`
		_, err = mysql.DB.Exec(sqlStr3, msg.PostID, msg.UserID, msg.Action)
		if err != nil {
			return err
		}
	}

	fmt.Println("接收到消息：", msg)
	return nil
}

func consumeMessage(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	for _, msg := range msgs {
		likeMsg := LikeMessage{}
		json.Unmarshal(msg.Body, &likeMsg)
		err := handleLikeEvent(likeMsg)
		if err != nil {
			return consumer.ConsumeRetryLater, err
		}
	}
	return consumer.ConsumeSuccess, nil
}

func main() {
	global.Consumer.Subscribe("vote_post",
		consumer.MessageSelector{},
		consumeMessage,
	)
	global.Consumer.Start()
	defer global.Consumer.Shutdown()
	time.Sleep(time.Hour * 24)
}
