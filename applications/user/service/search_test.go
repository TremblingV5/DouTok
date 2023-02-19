package service

import (
	"fmt"
	"github.com/TremblingV5/DouTok/applications/user/misc"
	"log"
	"testing"
	"time"
)

func TestSearchUser(t *testing.T) {
	Init()

	curr := fmt.Sprint(time.Now().Unix())

	userId, err, errNo := WriteNewUser(curr, curr)
	if err != nil {
		log.Panicln(err)
	} else if errNo != &misc.Success {
		log.Println(userId, errNo)
		log.Panicln(errNo)
	}

	u1, err1 := QueryUserByIdInRDB(userId)

	u2, err2 := QueryUserByIdInRDB(userId)

	u3, err3 := FindUserByUserName(curr)

	if err1 != nil || err2 != nil || err3 != nil {
		log.Panicln(err1, err2, err3)
	}

	log.Println(u1, u3)
	log.Println(u2)
}
