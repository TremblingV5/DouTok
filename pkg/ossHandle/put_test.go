package ossHandle

import (
	"fmt"
	"os"
	"testing"

	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/configurator"
)

func TestPut(t *testing.T) {
	file, _ := os.Open("/home/xinzf/Projects/OpenSource/DouTok/README.md")

	var config configStruct.OssConfig
	configurator.InitConfig(
		&config, "oss.yaml",
	)

	var client OssClient
	err := client.Init(
		config.Endpoint, config.Key,
		config.Secret, config.BucketName,
	)

	if err != nil {
		panic(err)
	}

	if err := client.Put(
		"video", "test.md", file, GetCallBackMap(config),
	); err != nil {
		fmt.Println(err)
	}
	defer file.Close()
}
