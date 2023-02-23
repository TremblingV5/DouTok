package ParamsValidator

import (
	"github.com/TremblingV5/DouTok/pkg/errno"
	"log"
	"testing"
)

var success = errno.NewErrNo(0, "Success")
var usernameErr = errno.NewErrNo(1, "usernameErr")
var passwordErr = errno.NewErrNo(2, "passwordErr")

func TestValidate(t *testing.T) {
	username := "xinzf"
	password := "123456"

	validator := New(&success)
	validator.Set(func() (bool, *errno.ErrNo) {
		if len(username) < 10 {
			return false, &usernameErr
		} else {
			return true, &success
		}
	}).Set(func() (bool, *errno.ErrNo) {
		if len(password) < 6 {
			return false, &passwordErr
		} else {
			return true, &success
		}
	})

	ok, errNo := validator.Validate()
	log.Println(ok)
	log.Println(errNo)
}
