package misc

import (
	"fmt"
	"testing"
)

type People struct {
	Key  string `json:"key"`
	Name string `json:"name"`
	Keys struct {
		Test  string `json:"test"`
		Fruit string `json:"fruit"`
	} `json:"keys"`
}

func TestMap2Struct(t *testing.T) {
	m := map[string]any{
		"key":  "value",
		"name": "Tom",
		"keys": map[string]string{
			"test":  "get",
			"fruit": "apple",
		},
	}

	p := People{}
	Map2Struct(m, &p)

	fmt.Println(p)
}
