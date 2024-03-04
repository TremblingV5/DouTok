package configStruct

type HBase struct {
	Host string `mapstructure:"Host" default:"localhost"`
}

type HBaseConfig struct {
	Host string `yaml:"host"`
}
