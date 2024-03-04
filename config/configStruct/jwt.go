package configStruct

type Jwt struct {
	SigningKey string `env:"JWT_SIGNING_KEY" envDefault:"signingKey" configPath:"JWT.signingKey"`
}
