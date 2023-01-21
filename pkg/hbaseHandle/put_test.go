package hbaseHandle

import (
	"testing"

	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/configurator"
	"github.com/sirupsen/logrus"
)

func TestPut(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	var config configStruct.HBaseConfig
	configurator.InitConfig(
		&config, "hbase.yaml",
	)

	client := InitHB(config.Host)

	values := map[string]map[string][]byte{
		"data": {
			"title":     []byte("test video"),
			"id":        []byte("0001"),
			"author_id": []byte("0002"),
			"video_url": []byte("www.baidu.com"),
			"cover_url": []byte("www.baidu.com"),
		},
	}

	client.Put("videos", "00010003", values)
}
