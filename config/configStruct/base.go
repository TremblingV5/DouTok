package configStruct

import "fmt"

type Base struct {
	Name    string `env:"SERVER_NAME" envDefault:"unknown"`
	Address string `env:"SERVER_ADDRESS" envDefault:"localhost"`
	Port    int    `env:"SERVER_PORT" envDefault:"8080"`
}

func (b Base) GetAddr() string {
	return fmt.Sprintf("%s:%d", b.Address, b.Port)
}

func (b Base) GetName() string {
	return fmt.Sprintf("%s", b.Name)
}
