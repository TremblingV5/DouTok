package configStruct

type MySQL struct {
	Host      string `env:"MYSQL_HOST" envDefault:"localhost" configPath:"MySQL.Host"`
	Port      int    `env:"MYSQL_PORT" envDefault:"3306" configPath:"MySQL.Port"`
	Username  string `env:"MYSQL_USERNAME" envDefault:"root" configPath:"MySQL.Username"`
	Password  string `env:"MYSQL_PASSWORD" envDefault:"root" configPath:"MySQL.Password"`
	Database  string `env:"MYSQL_DATABASE" envDefault:"DouTok" configPath:"MySQL.Database"`
	CharSet   string `env:"MYSQL_CHARSET" envDefault:"utf8mb4" configPath:"MySQL.CharSet"`
	ParseTime bool   `env:"MYSQL_PARSETIME" envDefault:"true" configPath:"MySQL.ParseTime"`
	Loc       string `env:"MYSQL_LOC" envDefault:"Local" configPath:"MySQL.loc"`
}

type MySQLConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
