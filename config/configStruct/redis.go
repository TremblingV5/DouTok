package configStruct

type Redis struct {
	Host     string
	Port     string
	Dsn      string `env:"REDIS_DSN" envDefault:"localhost:6379"`
	Password string `env:"REDIS_PASSWORD" envDefault:"root"`
	// {db name 1}:{db num 1},{db name 2}:{db num 2}
	Databases string `env:"REDIS_DATABASES" envDefault:""`
}

type RedisConfig struct {
	Host      string         `yaml:"host"`
	Port      string         `yaml:"port"`
	Password  string         `yaml:"password"`
	Databases map[string]int `yaml:"databases"`
}
