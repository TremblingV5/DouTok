package main

import (
	"log"

	"github.com/sirupsen/logrus"

	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/configurator"
	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	var config configStruct.HBaseConfig
	configurator.InitConfig(
		&config, "hbase.yaml",
	)

	client := hbaseHandle.InitHB(config.Host)

	values := map[string]map[string][]byte{
		"d": {
			"title": []byte("test video"),
			"time":  []byte("1998"),
		},
	}

	err := client.Put("test", "00030005", values)
	log.Println(err)
}
