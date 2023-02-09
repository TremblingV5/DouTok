package utils

import (
	"fmt"
	"testing"
)

func TestRedisKeyGen(t *testing.T) {
	r := KeyGen(123, 1, 1)
	fmt.Println(r)
	r = KeyGen(123, 1, 2)
	fmt.Println(r)
	r = KeyGen(123, 2, 1)
	fmt.Println(r)
	r = KeyGen(123, 2, 2)
	fmt.Println(r)
}
