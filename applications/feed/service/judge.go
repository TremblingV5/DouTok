package service

import "strconv"

/*
	以下不等式是否成立：
	t1 - t2 >= diff
*/
func JudgeTimeDiff(t1 int64, t2 string, diff int64) bool {
	t2_i, _ := strconv.Atoi(t2)
	t2_i64 := int64(t2_i)
	return t1-t2_i64 >= diff
}

/*
	以下不等式是否成立：
	q1 / q2 >= ratio
*/
func JudgeQuantityRatio(q1 float64, q2 float64, ratio float64) bool {
	return q1/q2 >= ratio
}
