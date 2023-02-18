package service

import (
	"fmt"
	"github.com/TremblingV5/DouTok/applications/user/misc"
	"log"
	"testing"
	"time"
)

func TestWriteNewUser(t *testing.T) {
	Init()

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
