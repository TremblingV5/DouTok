package mysqlIniter

import (
	"fmt"
	"testing"

	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/configurator"
)

func TestInitDb(t *testing.T) {
	var config configStruct.MySQLConfig
	configurator.InitConfig(
		&config, "mysql.yaml",
	)

	db, err := InitDb(
		config.Username, config.Password, config.Host, config.Port, config.Database,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(db)
}
