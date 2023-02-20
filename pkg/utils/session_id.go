package utils

import "fmt"

// 生成通讯会话id，待优化
func GenerateSessionId(a int64, b int64) string {
	if a < b {
		return fmt.Sprintf("%d%d", a, b)
	}
	return fmt.Sprintf("%d%d", b, a)
}
