package configurator

import (
	"os"

	"gopkg.in/yaml.v3"
)

func InitConfig(t any, configName string) error {
	configPath, err := GetConfigPath(configName)
	if err != nil {
		return err
	}

	file, _ := os.ReadFile(configPath)
	if err := yaml.Unmarshal(file, t); err != nil {
		return err
	}

	return nil
}
