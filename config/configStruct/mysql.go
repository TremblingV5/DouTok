package configStruct

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL struct {
	Host      string `env:"MYSQL_HOST" envDefault:"localhost" configPath:"MySQL.Host"`
	Port      int    `env:"MYSQL_PORT" envDefault:"3306" configPath:"MySQL.Port"`
	Username  string `env:"MYSQL_USERNAME" envDefault:"root" configPath:"MySQL.Username"`
	Password  string `env:"MYSQL_PASSWORD" envDefault:"root" configPath:"MySQL.Password"`
	Database  string `env:"MYSQL_DATABASE" envDefault:"DouTok" configPath:"MySQL.Database"`
	CharSet   string `env:"MYSQL_CHARSET" envDefault:"utf8mb4" configPath:"MySQL.CharSet"`
	ParseTime bool   `env:"MYSQL_PARSETIME" envDefault:"true" configPath:"MySQL.ParseTime"`
	Loc       string `env:"MYSQL_LOC" envDefault:"Local" configPath:"MySQL.loc"`
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
