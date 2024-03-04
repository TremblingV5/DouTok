package configStruct

import "fmt"

type Etcd struct {
	Address string `env:"ETCD_ADDRESS" envDefault:"localhost" configPath:"Etcd.Address"`
	Port    int    `env:"ETCD_PORT" envDefault:"2379" configPath:"Etcd.Port"`
}

func (e Etcd) GetAddr() string {
	return fmt.Sprintf("%s:%d", e.Address, e.Port)
}
