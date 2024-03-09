package configStruct

import (
	"github.com/minio/minio-go/v6"
)

type MinIO struct {
	Endpoint string `env:"MINIO_ENDPOINT" envDefault:"localhost:9000" configPath:"Minio.Endpoint"`
	Key      string `env:"MINIO_Key" envDefault:"root" configPath:"Minio.Key"`
	Secret   string `env:"MINIO_SECRET" envDefault:"rootroot" configPath:"Minio.Secret"`
	Bucket   string `env:"MINIO_BUCKET" envDefault:"DouTok" configPath:"Minio.Bucket"`
}

func (m *MinIO) InitIO() (*minio.Client, error) {
	return minio.New(
		m.Endpoint,
		m.Key,
		m.Secret,
		false,
	)
}
