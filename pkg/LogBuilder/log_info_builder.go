package LogBuilder

import (
	"fmt"
	"sort"
)

type LogInfoBuilder struct {
	logType string
	message string
	Items   map[string]string
}

func InitLogBuilder() *LogInfoBuilder {
	return &LogInfoBuilder{
		Items: make(map[string]string),
	}
}

func (b *LogInfoBuilder) SetLogType(typeString string) {
	b.logType = typeString
}

func (b *LogInfoBuilder) SetMessage(str string) {
	b.message = str
}

func (b *LogInfoBuilder) Collect(key string, value string) {
	tempKey := key

	_, ok := b.Items[tempKey]
	if ok {
		cnt := 1
		for {
			_, ok := b.Items[tempKey+fmt.Sprint(cnt)]
			if !ok {
				tempKey = tempKey + fmt.Sprint(cnt)
				break
			}
			cnt++
		}
	}

	b.Items[tempKey] = value
}

func (b *LogInfoBuilder) Get() []string {
	temp := []*logInfo{}

	for k, v := range b.Items {
		newLogInfo := logInfo{
			key:   k,
			value: v,
		}
		temp = append(temp, &newLogInfo)
	}

	sort.Sort(logInfoArray(temp))
	res := []string{}
	for _, v := range temp {
		res = append(res, v.key)
		res = append(res, v.value)
	}
	return res
}

func (b *LogInfoBuilder) Write(logger *Logger) {
	logger.Write(b.logType, b.message, b.Get()...)
}
