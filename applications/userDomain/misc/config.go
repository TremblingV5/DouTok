package misc

import "github.com/TremblingV5/DouTok/pkg/dtviper"

var Config *dtviper.Config

func InitViperConfig() {
	config := dtviper.ConfigInit(ViperConfigEnvPrefix, ViperConfigEnvFilename)
	Config = config
}

func GetConfig(key string) string {
	return Config.Viper.GetString(key)
}

func GetConfigNum(key string) int64 {
	return Config.Viper.GetInt64(key)
}
