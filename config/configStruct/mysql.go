package configStruct

type MySQL struct {
	Host      string `mapstructure:"Host" default:"localhost"`
	Port      int    `mapstructure:"Port" default:"3306"`
	Username  string `mapstructure:"Username" default:"root"`
	Password  string `mapstructure:"Password" default:"root"`
	Database  string `mapstructure:"Database" default:"DouTok"`
	CharSet   string `mapstructure:"CharSet" default:"utf8mb4"`
	ParseTime bool   `mapstructure:"ParseTime" default:"true"`
	Loc       string `mapstructure:"Loc" default:"Local"`
}

type MySQLConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
