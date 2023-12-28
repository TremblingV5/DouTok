package configStruct

type MySQL struct {
	Host      string `env:"MYSQL_HOST" envDefault:"localhost"`
	Port      int    `env:"MYSQL_PORT" envDefault:"3306"`
	Username  string `env:"MYSQL_USERNAME" envDefault:"root"`
	Password  string `env:"MYSQL_PASSWORD" envDefault:"root"`
	Database  string `env:"MYSQL_DATABASE" envDefault:"DouTok"`
	CharSet   string `env:"MYSQL_CHARSET" envDefault:"utf8mb4"`
	ParseTime bool   `env:"MYSQL_PARSETIME" envDefault:"true"`
	Loc       string `env:"MYSQL_LOC" envDefault:"Local"`
}

type MySQLConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
