package service

import (
	"encoding/json"
	"fmt"
)

func VideoList2String(list []VideoInHB) []string {
	res := []string{}

	for _, v := range list {
		r, err := json.Marshal(v)
		if err != nil {
			continue
		}
		res = append(res, string(r))
	}

	return res
}

func String2VideoList(list []string) []VideoInHB {
	res := []VideoInHB{}

	for _, v := range list {
		temp := VideoInHB{}
		err := json.Unmarshal([]byte(v), &temp)
		if err != nil {
			fmt.Println(err)
		}
		res = append(res, temp)
	}

	return res
}
