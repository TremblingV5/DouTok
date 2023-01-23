package ossHandle

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"mime/multipart"

	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func GetCallBackMap(config configStruct.OssConfig) string {
	callbackMap := map[string]string{}
	callbackMap["callbackUrl"] = config.Callback
	// callbackMap["callbackHost"] = config.Endpoint
	callbackMap["callbackBody"] = "filename=${object}&size=${size}&mimeType=${mimeType}"
	callbackMap["callbackBodyType"] = "application/x-www-form-urlencoded"

	callbackBuffer := bytes.NewBuffer([]byte{})
	callbackEncoder := json.NewEncoder(callbackBuffer)
	callbackEncoder.SetEscapeHTML(false)
	callbackEncoder.Encode(callbackMap)

	callbackVal := base64.StdEncoding.EncodeToString(callbackBuffer.Bytes())
	return callbackVal
}

func (o *OssClient) Put(objectType string, filename string, data multipart.File, callback string) error {
	err := o.Bucket.PutObject(
		objectType+"/"+filename,
		data,
		oss.Callback(callback),
	)

	if err != nil {
		return err
	}

	return nil
}