package ossHandle

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"

	"github.com/TremblingV5/DouTok/config/configStruct"
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
	if err := callbackEncoder.Encode(callbackMap); err != nil {
		return ""
	}

	callbackVal := base64.StdEncoding.EncodeToString(callbackBuffer.Bytes())
	return callbackVal
}

func (o *OssClient) Put(objectType string, filename string, data io.Reader) error {
	err := o.Bucket.PutObject(
		objectType+"/"+filename,
		data,
		// oss.Callback(callback),
	)

	if err != nil {
		return err
	}

	return nil
}
