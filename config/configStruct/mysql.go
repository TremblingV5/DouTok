package configStruct

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL struct {
	Host      string `mapstructure:"Host" default:"localhost"`
	Port      int    `mapstructure:"Port" default:"3306"`
	Username  string `mapstructure:"Username" default:"root"`
	Password  string `mapstructure:"Password" default:"root"`
	Database  string `mapstructure:"Database" default:"DouTok"`
	CharSet   string `mapstructure:"CharSet" default:"utf8mb4"`
	ParseTime bool   `mapstructure:"ParseTime" default:"true"`
	Loc       string `mapstructure:"Loc" default:"Local"`
}

func (m *MySQL) InitDB() (*gorm.DB, error) {
	Db, err := gorm.Open(
		mysql.Open(
			fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
				m.Username, m.Password, m.Host, m.Port, m.Database),
		),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	return Db, err
}

type MySQLConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
