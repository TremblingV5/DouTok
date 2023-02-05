package misc

import "testing"

type Obj struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestStruct2Map(t *testing.T) {
	stu := Obj{Name: "xiaoming", Age: 15}
	mp, err := Struct2Map(&stu)
	if err != nil {
		panic(err)
	}
	println("age = ", int64(mp["age"].(float64)))
	println("name = ", mp["name"].(string))
}
