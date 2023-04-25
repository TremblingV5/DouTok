package service

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func PasswordEncrypt(user_id int64, src string, salt string) string {
	opNum := GetOpNum(user_id)

	for i := 0; i < int(opNum); i++ {
		src = GetMd5(src + salt)
	}

	return src
}

/*
	生成salt
*/
func GenSalt() string {
	return GenRandString(32)
}

/*
	获取一个user id所对应的操作数
*/
func GetOpNum(id int64) int64 {
	str := fmt.Sprint(id)

	l := 0
	r := len(str) - 1
	var lNum, rNum string
	lNum = ""
	rNum = ""

	for {
		if string(str[l]) >= "0" && string(str[l]) <= "9" {
			lNum = string(str[l])
		} else {
			l++
		}

		if string(str[r]) >= "0" && string(str[r]) <= "9" {
			rNum = string(str[r])
		} else {
			r++
		}

		if l == r || (lNum != "" && rNum != "") {
			break
		}
	}

	if lNum == "" {
		lNum = "6"
	}

	if rNum == "" {
		rNum = "6"
	}

	res := lNum + rNum
	res_int, _ := strconv.Atoi(res)

	return int64(res_int)
}

/*
	进行md5加密
*/
func GetMd5(str string) string {
	code := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", code)
}

/*
	根据给定长度生成一个随机字符串
*/
func GenRandString(l int) string {
	list := []byte("0123456789abcdefghigklmnopqrstuvwxyz")

	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < l; i++ {
		result = append(result, list[r.Intn(len(list))])
	}

	return string(result)
}
