package configStruct

import "fmt"

type Otel struct {
	Host   string `env:"OTEL_HOST" envDefault:"localhost"`
	Port   int    `env:"OTEL_PORT" envDefault:"4317"`
	Enable bool
}

func (o Otel) GetAddr() string {
	return fmt.Sprintf("%s:%d", o.Host, o.Port)
}

func (o Otel) IsEnable() bool {
	return o.Enable
}
