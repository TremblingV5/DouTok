package hbaseHandle

import (
	"fmt"
	"testing"

	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/configurator"
)

func TestScan(t *testing.T) {
	var config configStruct.HBaseConfig
	configurator.InitConfig(
		&config, "hbase.yaml",
	)

	client := InitHB(config.Host)

	res, _ := client.Scan("test")

	for k, v := range res {
		fmt.Println(k)
		fmt.Println(v)
	}
}

func TestScanRange(t *testing.T) {
	var config configStruct.HBaseConfig
	configurator.InitConfig(
		&config, "hbase.yaml",
	)

	client := InitHB(config.Host)

	res, _ := client.ScanRange("test", "0000", "0003")

	for k, v := range res {
		fmt.Println(k)
		fmt.Println(v)
	}
}

func TestScanByRowKeyPrefix(t *testing.T) {
	var config configStruct.HBaseConfig
	configurator.InitConfig(
		&config, "hbase.yaml",
	)

	client := InitHB(config.Host)

	res, _ := client.Scan("test", GetFilterByRowKeyPrefix("0003")...)

	for k, v := range res {
		fmt.Println(k)
		fmt.Println(v)
	}
}

// TODO: 待修改，无法使用
func TestScanByRowKeyRange(t *testing.T) {
	var config configStruct.HBaseConfig
	configurator.InitConfig(
		&config, "hbase.yaml",
	)

	client := InitHB(config.Host)

	res, _ := client.Scan("test", GetFilterByRowKeyRange(100, "00010003", "00030003")...)

	for k, v := range res {
		fmt.Println(k)
		fmt.Println(v)
	}
}
