package ossHandle

import "github.com/aliyun/aliyun-oss-go-sdk/oss"

func (o *OssClient) Init(endpoint string, key string, secret string, bucketName string) error {
	client, err := oss.New(
		endpoint, key, secret,
	)

	if err != nil {
		return err
	}

	bucket, err := client.Bucket(bucketName)

	if err != nil {
		return err
	}

	o.Client = client
	o.Bucket = bucket

	return nil
}
