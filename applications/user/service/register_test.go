package service

import (
	"fmt"
	"github.com/TremblingV5/DouTok/applications/user/misc"
	"github.com/TremblingV5/DouTok/pkg/utils"
	"log"
	"testing"
	"time"
)

func TestWriteNewUser(t *testing.T) {
	misc.InitViperConfig()

	InitDb(
		misc.GetConfig("MySQL.Username"),
		misc.GetConfig("MySQL.Password"),
		misc.GetConfig("MySQL.Host"),
		misc.GetConfig("MySQL.Port"),
		misc.GetConfig("MySQL.Database"),
	)

	utils.InitSnowFlake(misc.GetConfigNum("Snowflake.Node"))

	curr := fmt.Sprint(time.Now().Unix())

	userId, err, errNo := WriteNewUser(curr, curr)
	if err != nil {
		log.Panicln(err)
	} else {
		log.Println(userId, errNo)
	}

	userId, err, errNo = WriteNewUser(curr, curr)
	if err != nil {
		log.Panicln(err)
	} else {
		if errNo != &misc.UserNameExistedErr {
			log.Panicln("插入重复用户未报错")
		}
	}
}
