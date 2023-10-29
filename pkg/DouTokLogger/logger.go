package DouTokLogger

import (
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/onsi/ginkgo/reporters/stenographer/support/go-isatty"
	"github.com/segmentio/ksuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

const (
	LogFormatDev    string = "dev"
	LogFormatNormal string = "normal"
	LogFormatSplunk string = "splunk"
)

func InitLogger(config configStruct.Logger) *zap.Logger {
	var loggerConfig zap.Config

	switch config.LogFormat {
	case LogFormatDev:
		loggerConfig = zap.NewDevelopmentConfig()
		if isatty.IsTerminal(os.Stdout.Fd()) {
			loggerConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
			loggerConfig.InitialFields = map[string]interface{}{
				"session": ksuid.New(),
			}
		}
	default:
		loggerConfig = zap.NewProductionConfig()

		if config.Level == 0 {
			loggerConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
		} else {
			loggerConfig.Level = zap.NewAtomicLevelAt(config.Level)
		}

		loggerConfig.Development = config.Development
		loggerConfig.Encoding = config.Encoding
		loggerConfig.EncoderConfig.MessageKey = config.EncoderConfig.MessageKey
		loggerConfig.EncoderConfig.LevelKey = config.EncoderConfig.LevelKey
		loggerConfig.EncoderConfig.TimeKey = config.EncoderConfig.TimeKey
		loggerConfig.EncoderConfig.NameKey = config.EncoderConfig.NameKey
		loggerConfig.EncoderConfig.CallerKey = config.EncoderConfig.CallerKey
		loggerConfig.EncoderConfig.FunctionKey = config.EncoderConfig.FunctionKey
		loggerConfig.EncoderConfig.StacktraceKey = config.EncoderConfig.StacktraceKey
		loggerConfig.EncoderConfig.SkipLineEnding = config.EncoderConfig.SkipLineEnding
		loggerConfig.EncoderConfig.LineEnding = config.EncoderConfig.LineEnding
		loggerConfig.EncoderConfig.EncodeLevel = func() zapcore.LevelEncoder {
			switch config.EncoderConfig.LevelEncoder {
			case "capitalColor":
				return zapcore.CapitalColorLevelEncoder
			case "lowercase":
				return zapcore.LowercaseLevelEncoder
			default:
				return zapcore.CapitalLevelEncoder
			}
		}()
		loggerConfig.EncoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(time.Format("2006-01-02 15:04:05.000"))
		}
		loggerConfig.EncoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
		loggerConfig.EncoderConfig.EncodeCaller = func() zapcore.CallerEncoder {
			switch config.EncoderConfig.CallerEncoder {
			case "short":
				return zapcore.ShortCallerEncoder
			default:
				return zapcore.FullCallerEncoder
			}
		}()
		loggerConfig.EncoderConfig.EncodeName = func() zapcore.NameEncoder {
			switch config.EncoderConfig.NameEncoder {
			case "short":
				return zapcore.FullNameEncoder
			default:
				return zapcore.FullNameEncoder
			}
		}()
		loggerConfig.EncoderConfig.ConsoleSeparator = config.EncoderConfig.ConsoleSeparator
		loggerConfig.OutputPaths = config.OutputPaths
	}

	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}

	return logger
}
