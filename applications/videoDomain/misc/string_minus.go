package misc

import (
	"fmt"
	"strconv"
)

func TimestampMinus(curr string, m int) string {
	currTime, _ := strconv.Atoi(curr)
	return fmt.Sprint(currTime - int(m))
}
