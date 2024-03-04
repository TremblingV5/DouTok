package configStruct

import "fmt"

type Etcd struct {
	Address string `mapstructure:"Address" default:"localhost"`
	Port    int    `mapstructure:"Port" default:"2379"`
}

func (e Etcd) GetAddr() string {
	return fmt.Sprintf("%s:%d", e.Address, e.Port)
}
