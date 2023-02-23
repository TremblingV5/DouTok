package LogBuilder

type logInfoArray []*logInfo

func (arr logInfoArray) Len() int {
	return len(arr)
}

func (arr logInfoArray) Less(i, j int) bool {
	return arr[i].key < arr[j].key
}

func (arr logInfoArray) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
