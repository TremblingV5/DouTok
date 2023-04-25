package minioHandle

import "github.com/minio/minio-go/v6"

func (c *MinioClient) Init(endpoint string, key string, secret string, bucketName string) {
	client, err := minio.New(
		endpoint, key, secret, false,
	)

	if err != nil {
		panic(err)
	}

	c.Client = client
	c.Bucket = bucketName
}
