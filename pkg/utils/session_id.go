package utils

// 生成通讯会话id，待优化
func GenerateSessionId(a int64, b int64) string {
	if a < b {
		return string(a) + string(b)
	}
	return string(b) + string(a)
}
