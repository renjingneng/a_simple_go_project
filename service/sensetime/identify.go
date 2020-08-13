// @Description  对外提供方法
// @Author  renjingneng
// @CreateTime  2020/8/10 21:01
package sensetime

import (
	"errors"
	"fmt"
	"time"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"

	"github.com/renjingneng/a_simple_go_project/lib/httplib"
)

var apiKey string
var apiSecret string

// Init 初始化
//
// @Author  renjingneng
//
// @CreateTime  2020/8/13 15:57
func Init(Key, Secret string) {
	apiKey = Key
	apiSecret = Secret
}

// OcrIdcard 检查身份证
//
// @Author  renjingneng
//
// @CreateTime  2020/8/10 21:01
func OcrIdcard(localFilePath string, side string) (string, error) {
	if err := preProcess(); err != nil {
		return "", err
	}
	data := map[string]string{
		"auto_rotate":  "true",
		"return_score": "true",
		"classify":     "true",
		"side":         side,
		"image_file":   localFilePath,
	}
	url := "https://v2-auth-api.visioncloudapi.com/ocr/idcard/stateless"
	str, err := publicPost(data, url, "image_file")
	return postProcess(str, err)
}

func preProcess() error {
	if apiKey == "" {
		return errors.New("apiKey参数为空")
	}
	if apiSecret == "" {
		return errors.New("apiSecret参数为空")
	}
	return nil
}
func postProcess(str string, err error) (string, error) {
	if err != nil {
		return str, err
	}
	code := gjson.Get(str, "code").String()
	livenessStatus := gjson.Get(str, "liveness_status").String()
	if code != "" {
		str, _ = sjson.Set(str, "code_info", codeMap[code])
	}
	if livenessStatus != "" {
		str, _ = sjson.Set(str, "liveness_status_cn", livenessStatusMap[livenessStatus])
	}
	return str, err
}

func publicPost(data map[string]string, url string, fileName string) (string, error) {
	nonce := makeNonce(16)
	timestamp := fmt.Sprint(time.Now().Unix() * 1000)
	stringSignature := makeStringSignature(nonce, timestamp, apiKey)
	signature := signString(stringSignature, apiSecret)
	Authorization := "key=" + apiKey + ",timestamp=" + timestamp + ",nonce=" + nonce + ",signature=" + signature
	http := httplib.Post(url)
	http.Header("Content-Type", "application/json")
	http.Header("Authorization", Authorization)
	for i, v := range data {
		if i == fileName {
			http.PostFile(i, v)
		} else {
			http.Param(i, v)
		}
	}
	str, err := http.String()
	return str, err
}
