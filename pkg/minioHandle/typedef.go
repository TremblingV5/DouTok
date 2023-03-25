package minioHandle

import "github.com/minio/minio-go/v6"

type MinioClient struct {
	Client *minio.Client
	Bucket string
}
