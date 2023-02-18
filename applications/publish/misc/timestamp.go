package misc

import "fmt"

var MAX_TIMESTAMP = 99999999999

func GetTimeRebound(timestamp int64) string {
	return fmt.Sprint(int64(MAX_TIMESTAMP) - timestamp)
}
