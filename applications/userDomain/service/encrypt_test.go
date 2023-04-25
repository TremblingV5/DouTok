package service

import "testing"

func TestPasswordEncrypt(t *testing.T) {
	user_id := 1678546894123654781
	src := "DouTokNo1@"
	salt := GenSalt()

	password1 := PasswordEncrypt(int64(user_id), src, salt)
	password2 := PasswordEncrypt(int64(user_id), src, salt)

	if password1 != password2 {
		panic("加密方案不能保证结果一致")
	}
}
