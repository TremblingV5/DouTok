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
		"d": {
			"title": []byte("test video"),
			"time":  []byte("1998"),
		},
	}

	client.Put("test", "00030005", values)
}
