package configStruct

import "fmt"

type Base struct {
	Name     string `env:"SERVER_NAME" envDefault:"unknown" configPath:"Server.Name"`
	Address  string `env:"SERVER_ADDRESS" envDefault:"localhost" configPath:"Server.Address"`
	Port     int    `env:"SERVER_PORT" envDefault:"8080" configPath:"Server.Port"`
	NameCode int32  `env:"NAME_CODE" envDefault:"0" configPath:"Server.NameCode"`
	NodeCode int32  `env:"NODE_CODE" envDefault:"0" configPath:"Server.NodeCode"`
}

func (b Base) GetAddr() string {
	return fmt.Sprintf("%s:%d", b.Address, b.Port)
}

// TODO: GetName should be add more details about how to confirm a service name
func (b Base) GetName() string {
	return b.Name
}

func (b Base) GetNameCode() int32 {
	return b.NameCode
}

func (b Base) GetNodeCode() int32 {
	return b.NodeCode
}
