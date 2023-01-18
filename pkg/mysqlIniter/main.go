package mysqlIniter

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb(username string, password string, host string, port string, db string) (*gorm.DB, error) {
	Db, err := gorm.Open(
		mysql.Open(
			username+":"+password+"@tcp("+host+":"+port+")/"+db+"?charset=utf8&parseTime=True&loc=Local",
		),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)

	return Db, err
}
