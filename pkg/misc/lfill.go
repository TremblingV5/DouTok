package misc

func LFill(num string, bit int) string {
	if len(num) >= bit {
		return num
	} else {
		more := ""

		for i := 0; i < bit-len(num); i++ {
			more += "0"
		}

		return more + num
	}
}
