package main

import (
	"context"
	"fmt"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"github.com/TremblingV5/DouTok/pkg/kafka"
)

func main() {

	cfg := dtviper.ConfigInit("DOUTOK_MESSAGE", "message")

	cGroup := kafka.InitConsumerGroup(cfg.Viper.GetStringSlice("Kafka.Brokers"), cfg.Viper.GetString("Kafka.GroupId"))

	for {
		err := cGroup.Consume(context.Background(), cfg.Viper.GetStringSlice("Kafka.Topics"), kafka.ConsumerGroup)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}

	_ = cGroup.Close()
}
