package service

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/TremblingV5/DouTok/applications/userDomain/misc"
)

func TestWriteNewUser(t *testing.T) {
	Init()

	curr := fmt.Sprint(time.Now().Unix())

	userId, err, errNo := NewWriteNewUserService(context.Background()).WriteNewUser(curr, curr)
	if err != nil {
		log.Panicln(err)
	} else {
		log.Println(userId, errNo)
	}

	userId, err, errNo = NewWriteNewUserService(context.Background()).WriteNewUser(curr, curr)
	if err != nil {
		log.Panicln(err)
	} else {
		if errNo != &misc.UserNameExistedErr {
			log.Panicln("插入重复用户未报错")
		}
	}
}
