package configStruct

type Redis struct {
	Dsn      string `env:"REDIS_DSN" envDefault:"localhost:6379" configPath:"Redis.Dest"`
	Password string `env:"REDIS_PASSWORD" envDefault:"root" configPath:"Redis.Password"`
	// {db name 1}:{db num 1},{db name 2}:{db num 2}
	Databases string `mapstructure:"Databases" default:""`
}
