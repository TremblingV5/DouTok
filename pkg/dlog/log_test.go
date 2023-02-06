package dlog

import (
	"testing"
)

func TestHzLog(t *testing.T) {
	logger := InitHertzLog(3)
	logger.Info("hz test log1")
	logger.Info("hz test log2")
	logger.Info("hz test log3")
	logger.Info("hz test log4")
}

func TestKitexLog(t *testing.T) {
	logger := InitLog(3)
	logger.Info("kitex test log1")
	logger.Info("kitex test log2")
	logger.Info("kitex test log3")
	logger.Info("kitex test log4")
}
