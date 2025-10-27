package global

import (
	"context"
	"fmt"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"go.uber.org/zap"
)

var Consumer rocketmq.PushConsumer

func InitConsumer(nameServer []string, groupName, topic string) {
	// 创建 PushConsumer（推荐）
	var err error
	Consumer, err = rocketmq.NewPushConsumer(
		consumer.WithNameServer(nameServer), // NameServer 地址
		consumer.WithGroupName(groupName),   // 消费者组名称
		consumer.WithConsumeFromWhere(consumer.ConsumeFromLastOffset),
	)
	if err != nil {
		zap.S().Fatalf("创建消费者失败: %s", err.Error())
	}
	// 订阅主题 + 回调函数
	Consumer.Subscribe(topic, consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, msg := range msgs {
			fmt.Printf("接收到消息：%s  内容：%s\n", msg.MsgId, string(msg.Body))
		}
		return consumer.ConsumeSuccess, nil
	})
	if err := Consumer.Start(); err != nil {
		zap.S().Fatalf("启动消费者失败: %s", err.Error())
	}
}

func CloseConsumer() {
	if err := Consumer.Shutdown(); err != nil {
		zap.S().Fatalf("关闭消费者失败: %s", err.Error())
	}
}
