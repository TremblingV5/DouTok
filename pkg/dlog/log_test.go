package dlog

import (
	"testing"
)

func TestHzLog(t *testing.T) {
	logger := InitHertzLog(3)
	logger.Info("hz integration log1")
	logger.Info("hz integration log2")
	logger.Info("hz integration log3")
	logger.Info("hz integration log4")
}

func TestKitexLog(t *testing.T) {
	logger := InitLog(3)
	logger.Info("kitex integration log1")
	logger.Info("kitex integration log2")
	logger.Info("kitex integration log3")
	logger.Info("kitex integration log4")
}
