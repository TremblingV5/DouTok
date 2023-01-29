package hbaseHandle

import (
	"fmt"
	"testing"

	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/configurator"
)

func TestInitHB(t *testing.T) {
	var config configStruct.HBaseConfig
	configurator.InitConfig(
		&config, "hbase.yaml",
	)

	client := InitHB(config.Host)
	fmt.Println(client)
}
