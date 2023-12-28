package configStruct

type Snowflake struct {
	Node int64 `env:"SNOWFLAKE_NODE" envDefault:"1"`
}
