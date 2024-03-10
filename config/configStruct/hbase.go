package configStruct

import "github.com/tsuna/gohbase"

type HBase struct {
	Host string `env:"HBASE_HOST" envDefault:"localhost" configPath:"HBase.Host"`
}

func (h *HBase) InitHB() *gohbase.Client {
	c := gohbase.NewClient(h.Host)
	return &c
}

type HBaseConfig struct {
	Host string `yaml:"host"`
}
