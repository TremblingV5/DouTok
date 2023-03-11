package LogBuilder

import (
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	Handle *zap.Logger
}

type logInfo struct {
	key   string
	value string
}

func NewByViper(config *dtviper.Config) *Logger {
	filename := config.Viper.GetString("Log.path")
	maxSize := config.Viper.GetInt("Log.maxSize")
	maxBackups := config.Viper.GetInt("Log.maxBackups")
	maxAge := config.Viper.GetInt("Log.maxAge")
	skip := config.Viper.GetInt("Log.skip")

	return New(filename, maxSize, maxBackups, maxAge, skip)
}

func New(filename string, maxSize int, maxBackups int, maxAge int, skip int) *Logger {
	writeSyncer := getLogWriter(
		filename, maxSize, maxBackups, maxAge,
	)

	encoder := getEncoder()

	var l = new(zapcore.Level)
	if err := l.UnmarshalText([]byte("info")); err != nil {
		panic(err)
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(skip))
	return &Logger{
		Handle: logger,
	}
}

func (l *Logger) Write(logType string, msg string, values ...string) {
	if len(values)%2 != 0 {
		panic("日志信息的键与值数量不匹配")
	}

	logInfo := make([]zapcore.Field, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		logInfo[i/2] = zap.String(values[i], values[i+1])
	}

	switch logType {
	case "info":
		l.Handle.Info(msg, logInfo...)
	case "error":
		l.Handle.Error(msg, logInfo...)
	case "warn":
		l.Handle.Warn(msg, logInfo...)
	default:
		l.Handle.Info(msg, logInfo...)
	}
}
