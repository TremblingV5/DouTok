package configStruct

type OssConfig struct {
	Endpoint   string `yaml:"endpoint"`
	Key        string `yaml:"key"`
	Secret     string `yaml:"secret"`
	BucketName string `yaml:"bucket"`
}
