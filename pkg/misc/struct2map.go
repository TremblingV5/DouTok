package misc

import "encoding/json"

func Struct2Map(s any) (map[string]interface{}, error) {
	str, err := json.Marshal(s)

	if err != nil {
		return nil, err
	}

	var m map[string]interface{}
	if err = json.Unmarshal([]byte(str), &m); err != nil {
		return nil, err
	}

	return m, nil
}
