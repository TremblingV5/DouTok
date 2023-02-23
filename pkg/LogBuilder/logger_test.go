package LogBuilder

import "testing"

func TestLogger(t *testing.T) {
	logger := New("./temp.log", 1024*1024, 3, 10)

	log := InitLogBuilder()
	log.SetLogType("error")
	log.SetMessage("Test log info")
	log.Collect("key1", "value1")
	log.Collect("key2", "value2")
	log.Collect("key2", "value2")
	log.Collect("key3", "value3")

	log.Write(logger)
}
