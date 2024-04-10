package configStruct

type RedisConfig struct {
	Host      string         `yaml:"host"`
	Port      string         `yaml:"port"`
	Password  string         `yaml:"password"`
	Databases map[string]int `yaml:"databases"`
}
