package misc

import (
	"fmt"
	"strconv"
)

func TimestampMinus(curr string, m int) string {
	curr_time, _ := strconv.Atoi(curr)
	return fmt.Sprint(curr_time - int(m))
}
