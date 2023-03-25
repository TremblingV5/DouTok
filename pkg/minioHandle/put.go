package minioHandle

import (
	"io"

	"github.com/minio/minio-go/v6"
)

func (c *MinioClient) Put(objectType string, filename string, data io.Reader, size int) error {
	_, err := c.Client.PutObject(
		c.Bucket,
		objectType+"/"+filename,
		data,
		int64(size),
		minio.PutObjectOptions{},
	)

	if err != nil {
		return err
	}

	return nil
}
