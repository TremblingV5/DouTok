package misc

import "github.com/TremblingV5/DouTok/pkg/LogBuilder"

var Logger *LogBuilder.Logger

func InitLogger() {
	Logger = LogBuilder.New("./tmp/videoDomain.log", 1024*1024, 3, 10)
}
