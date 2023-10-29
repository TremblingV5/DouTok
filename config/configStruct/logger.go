package configStruct

import "go.uber.org/zap/zapcore"

type Logger struct {
	LogFormat     string        `env:"LOGGER_FORMAT" envDefault:"dev"` // json, console
	Level         zapcore.Level `env:"LOGGER_LEVEL" envDefault:"info"`
	Development   bool          `env:"LOGGER_DEVELOPMENT" envDefault:"false"`
	Encoding      string        `env:"LOGGER_ENCODING" envDefault:"json"`
	EncoderConfig struct {
		MessageKey       string `env:"MESSAGE_KEY" envDefault:"message"`
		LevelKey         string `env:"LEVEL_KEY" envDefault:"level"`
		TimeKey          string `env:"TIME_KEY" envDefault:"ts"`
		NameKey          string `env:"NAME_KEY" envDefault:"DouTokLogger"`
		CallerKey        string `env:"CALLER_KEY" envDefault:"caller"`
		FunctionKey      string `env:"FUNCTION_KEY" envDefault:"function"`
		StacktraceKey    string `env:"STACKTRACE_KEY" envDefault:"stacktrace"`
		SkipLineEnding   bool   `env:"SKIP_LINE_ENDING" envDefault:"false"`
		LineEnding       string `env:"LINE_ENDING" envDefault:"\n"`
		LevelEncoder     string `env:"LEVEL_ENCODER" envDefault:"capital"`    // capitalColor, capital, color, lowercase
		DurationEncoder  string `env:"DURATION_ENCODER" envDefault:"seconds"` // string, nanos, ms, seconds
		CallerEncoder    string `env:"CALLER_ENCODER" envDefault:"full"`      // short, full
		NameEncoder      string `env:"NAME_ENCODER" envDefault:"full"`        // short, full
		ConsoleSeparator string `env:"CONSOLE_SEPARATOR" envDefault:" "`
	} `envPrefix:"LOGGER_ENCODER_"`
	OutputPaths []string `env:"LOGGER_OUTPUT_PATHS" envDefault:"stdout"`
}
