package service

import "testing"

func Init() {
	InitViper()
	InitHB()
	InitRedisClient()
	InitSyncProducer()
	InitConsumerGroup()
	InitId()
}

func TestInit(t *testing.T) {
	Init()
}
