package ossHandle

import (
	"fmt"
	"testing"

	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/configurator"
)

func TestInit(t *testing.T) {
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
		fmt.Println(err)
	}

	fmt.Println(client.Bucket)
}
