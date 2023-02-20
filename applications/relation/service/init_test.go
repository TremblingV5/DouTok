package service

import "testing"

func Init() {
	InitViper()
	InitRedisClient()
	InitSyncProducer()
	InitConsumerGroup()
	InitId()
	InitDB()
	InitSafeMap()
	InitMutex()
	// TODO 待测试 Flush 协程
}

func TestInit(t *testing.T) {
	Init()
}
