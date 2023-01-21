package ossHandle

import "github.com/aliyun/aliyun-oss-go-sdk/oss"

type OssClient struct {
	Client *oss.Client
	Bucket *oss.Bucket
}
