package misc

import (
	"encoding/json"
)

func Map2Struct(m map[string]interface{}, s any) error {
	str, err := json.Marshal(m)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(str, s); err != nil {
		return err
	}

	return nil
}

func Map2Struct4HB(m map[string][]byte, s any) error {
	str, err := json.Marshal(m)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(str, s); err != nil {
		return err
	}

	return nil
}
