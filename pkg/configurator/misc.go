package configurator

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/TremblingV5/DouTok/pkg/constants"
)

func GetConfigPath(configName string) (string, error) {
	pathList := [5]string{
		"./config/",
		"../../config/",
		"../../../config/",
		"../../../../config/",
		"../../../../../config/",
	}

	for i := range pathList {
		_, err := os.Stat(pathList[i] + configName)
		if err == nil {
			p, _ := filepath.Abs(pathList[i] + configName)
			return p, nil
		}
	}

	return "", errors.New(constants.ErrConfigFileNotFound + ", file name: " + configName)
}
