package configStruct

type Snowflake struct {
	Node int64 `mapstructure:"SNOWFLAKE_NODE" default:"1"`
}
