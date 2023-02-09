package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var UserConfig, MessageConfig, RelationConfig *viper.Viper

func InitConfig(path string) {
	UserConfig, _ = getConfig(path, "user")
	MessageConfig, _ = getConfig(path, "message")
	RelationConfig, _ = getConfig(path, "relation")
}

func getConfig(path, name string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(name)
	v.AddConfigPath(path)
	err := v.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config changed!")
	})
	v.WatchConfig()
	return v, nil
}
