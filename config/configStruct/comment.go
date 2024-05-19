package configStruct

type CommentConfig struct {
	Global struct {
		Source   string `yaml:"Source"`
		ChangeMe string `yaml:"ChangeMe"`
	} `yaml:"Global"`
	JWT struct {
		SigningKey string `yaml:"signingKey"`
	} `yaml:"JWT"`
	Etcd struct {
		Address string `yaml:"Address"`
		Port    string `yaml:"Port"`
	} `yaml:"Etcd"`
	Server struct {
		Name    string `yaml:"Name"`
		Address string `yaml:"Address"`
		Port    string `yaml:"Port"`
	} `yaml:"Server"`
	Client struct {
		Echo    bool     `yaml:"Echo"`
		Foo     string   `yaml:"Foo"`
		Servers []string `yaml:"Servers"`
	} `yaml:"Client"`
	MySQLConfig struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Database string `yaml:"database"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"MySQL"`
}
