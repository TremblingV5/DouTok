package configStruct

type HBase struct {
	Host string `env:"HBASE_HOST" envDefault:"localhost"`
}

type HBaseConfig struct {
	Host string `yaml:"host"`
}
