package configStruct

import "fmt"

type Base struct {
	Name     string `mapstructure:"Name" default:"unknown" `
	Address  string `mapstructure:"Address" default:"localhost"`
	Port     int    `mapstructure:"Port" default:"8080"`
	NameCode int32  `mapstructure:"NAME_CODE" default:"0"`
	NodeCode int32  `mapstructure:"NODE_CODE" default:"0"`
}

func (b Base) GetAddr() string {
	return fmt.Sprintf("%s:%d", b.Address, b.Port)
}

func (b Base) GetName() string {
	return fmt.Sprintf("%s", b.Name)
}

func (b Base) GetNameCode() int32 {
	return b.NameCode
}

func (b Base) GetNodeCode() int32 {
	return b.NodeCode
}
