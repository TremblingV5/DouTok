package configStruct

type Redis struct {
	Host     string `mapstructure:"Host" default:"localhost"`
	Port     string `mapstructure:"Port" default:"6379"`
	Dsn      string `mapstructure:"Dsn" default:"localhost:6379"`
	Password string `mapstructure:"Password" default:"root"`
	// {db name 1}:{db num 1},{db name 2}:{db num 2}
	Databases string `mapstructure:"Databases" default:""`
}

type RedisConfig struct {
	Host      string         `yaml:"host"`
	Port      string         `yaml:"port"`
	Password  string         `yaml:"password"`
	Databases map[string]int `yaml:"databases"`
}
