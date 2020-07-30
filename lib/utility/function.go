package utility

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"net/url"
	"sort"
)

func ApiFormatData(errorCode int, data interface{}, ctx iris.Context) {
	r := make(map[string]interface{})

	// 判断map是否存在key
	if errorCode != 0 {
		r["error_code"] = errorCode

		if msg, ok := ErrorCodeMap[errorCode]; !ok {
			r["error"] = data
		} else {
			r["error"] = msg
		}
	} else {
		r["entry"] = data
	}

	callback := ctx.URLParam("callback")
	if callback == "" {
		ctx.JSON(r)
	} else {
		options := context.JSONP{"", callback}
		ctx.JSONP(r, options)
	}
}

func ParsingXcxParams(ctx iris.Context) (map[string]interface{}, map[string]string) {
	req_data := ctx.FormValue("business_param")
	if req_data == "" {
		return map[string]interface{}{}, map[string]string{}
	}
	oauth_info := map[string]string{
		"wa_code": ctx.FormValue("oauth_info[wa_code]"),
		"open_id": ctx.FormValue("oauth_info[open_id]"),
	}

	// json decode
	req_params := map[string]interface{}{}
	err := json.Unmarshal([]byte(req_data), &req_params)
	if err != nil {
		fmt.Println(err)
	}

	return req_params, oauth_info
}

func CreateSign(params map[string]string, token string) string {

	// 因为map是无序的所以将所有的key拿到排序
	keys := make([]string, len(params))
	i := 0
	for k, _ := range params {
		keys[i] = k
		i++
	}

	// 创建url.Vaules结构体变量
	u := url.Values{}

	// 对字符串slice进行排序
	sort.Strings(keys)

	// 根据顺序获取参数
	for _, k := range keys {
		u.Add(k, params[k])
	}
	// 把参数URL Encode
	str := u.Encode()

	// 参数拼接
	new := str + token
	// md5加密
	has := md5.Sum([]byte(new))
	// 转成换进制
	md5str := fmt.Sprintf("%x", has)

	return md5str
}

/**
签名方法
@author yongliang2@leju.com
@return string
*/
func CreateNewdelSign(params map[string]string, token string) string {

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
