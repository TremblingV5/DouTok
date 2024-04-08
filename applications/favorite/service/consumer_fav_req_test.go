package service

import "testing"

func TestConsumer4UpdateCount(t *testing.T) {
	Init()

	go Consumer4UpdateCount()
}
