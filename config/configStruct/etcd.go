package configStruct

import "fmt"

type Etcd struct {
	Address string `env:"ETCD_ADDRESS" envDefault:"localhost"`
	Port    int    `env:"ETCD_PORT" envDefault:"2379"`
}

func (e Etcd) GetAddr() string {
	return fmt.Sprintf("%s:%d", e.Address, e.Port)
}
