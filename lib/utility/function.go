package utility

import (
	"crypto/md5"
	"fmt"
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"sort"
	"time"

	"github.com/renjingneng/a_simple_go_project/core/config"
)

func MakeNonce(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	nonce := ""
	l := len(chars)
	rand.Seed(time.Now().Unix())
	for i := 0; i < length; i++ {
		r := rand.Intn(l)
		nonce = nonce + string(chars[r])
	}
	return nonce
}
func MoveUploadedFile(file multipart.File, path string) {
	out, _ := os.OpenFile(path,
		os.O_WRONLY|os.O_CREATE, 0666)
	io.Copy(out, file)
}

// GenerateUploadPath 生成上传路径，XX/upload/pathType/year/month/day/
//
// @Author  renjingneng
//
// @CreateTime  2020/8/14 11:07
func GenerateUploadPath(pathType string) string {
	year := fmt.Sprint(time.Now().Year())
	month := fmt.Sprint(time.Now().Month())
	day := fmt.Sprint(time.Now().Day())
	var path string
	if pathType == "pic" {
		path = config.Config["UploadPathPic"] + year + "/" + month + "/" + day + "/"
	} else if pathType == "video" {
		path = config.Config["UploadPathVideo"] + year + "/" + month + "/" + day + "/"
	} else {
		path = config.Config["UploadPathTemp"] + year + "/" + month + "/" + day + "/"
	}
	os.MkdirAll(path, 0666)
	return path
}

// ImageConfig 由于DecodeConfig方法调用需要引入太多包，所以为其包装一个新的方法
//
// @Author  renjingneng
//
// @Return
//(image.Config) {
//  ColorModel: (*color.modelFunc)(0xc0000062e0)({
//  f: (func(color.Color) color.Color) 0xb5f400
//  }),
//  Width: (int) 887,
//  Height: (int) 1920
//}
//(string) (len=4) "tiff"
//(error) <nil>
// @CreateTime  2020/8/14 15:27
func ImageConfig(r io.Reader) (image.Config, string, error) {
	return image.DecodeConfig(r)
}

/**
签名方法
@author yongliang2@leju.com
@return string
*/
func CreateSign(params map[string]string, token string) string {

	// 因为map是无序的所以将所有的key拿到排序
	keys := make([]string, len(params))
	i := 0
	for k, _ := range params {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	str := ""
	for _, v := range keys {
		if str == "" {
			str += v + "=" + params[v]
		} else {
			str += "&" + v + "=" + params[v]
		}
	}

	// 参数拼接
	new := str + token
	// md5加密
	has := md5.Sum([]byte(new))
	// 转成换进制
	md5str := fmt.Sprintf("%x", has)

	return md5str
}
