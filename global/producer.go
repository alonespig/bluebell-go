package global

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"go.uber.org/zap"
)

var Producer rocketmq.Producer

func InitProducer() {

	p, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{"172.27.49.67:9876"}),
		producer.WithGroupName("vote_post"),
	)
	if err != nil {
		zap.S().Fatalln("failed to create producer", err)
	}
	err = p.Start()
	if err != nil {
		zap.S().Fatalln("failed to start producer", err)
	}
	Producer = p
}
