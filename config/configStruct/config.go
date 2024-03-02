package configStruct

import (
	"github.com/TremblingV5/DouTok/pkg/dtviper"
)

type Config struct {
	Base      Base      `envPrefix:"DOUTOK_USER_"`
	Etcd      Etcd      `envPrefix:"DOUTOK_USER_"`
	Jwt       Jwt       `envPrefix:"DOUTOK_USER_"`
	MySQL     MySQL     `envPrefix:"DOUTOK_USER_"`
	Snowflake Snowflake `envPrefix:"DOUTOK_USER_"`
	HBase     HBase     `envPrefix:"DOUTOK_USER_"`
	Redis     Redis     `envPrefix:"DOUTOK_USER_"`
	Otel      Otel      `envPrefix:"DOUTOK_USER_"`
	Logger    Logger    `envPrefix:"DOUTOK_USER_"`
	viper     *dtviper.Config
}

func (config *Config) InitViper(envPrefix string, server string) {
	config.viper = dtviper.ConfigInit(envPrefix, server)
}

func (config *Config) ResolveViperConfig() {
	v := config.viper.Viper

	// Base
	config.Base.Name = v.GetString("Server.Name")
	config.Base.Address = v.GetString("Server.Address")
	config.Base.Port = v.GetInt("Server.Port")

	// Etcd
	config.Etcd.Address = v.GetString("Etcd.Address")
	config.Etcd.Port = v.GetInt("Etcd.Port")

	// Mysql
	config.MySQL.Host = v.GetString("MySQL.Host")
	config.MySQL.Port = v.GetInt("MySQL.Port")
	config.MySQL.Username = v.GetString("MySQL.Username")
	config.MySQL.Password = v.GetString("MySQL.Password")
	config.MySQL.Database = v.GetString("MySQL.Database")

	// HBae
	config.HBase.Host = v.GetString("HBase.Host")

	// Redis
	config.Redis.Host = v.GetString("Redis.Host")
	config.Redis.Port = v.GetString("Redis.Port")
	config.Redis.Password = v.GetString("Redis.Password")
	config.Redis.Databases = v.GetString("Redis.DataBases.Default")

	// Otel
	config.Otel.Enable = v.GetBool("Otel.Enable")
	config.Otel.Host = v.GetString("Otel.Host")
	config.Otel.Port = v.GetInt("Otel.Port")

}
