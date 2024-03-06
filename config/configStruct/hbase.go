package configStruct

type HBase struct {
	Host string `env:"HBASE_HOST" envDefault:"localhost" configPath:"Hbase:Host"`
}

type HBaseConfig struct {
	Host string `yaml:"host"`
}
