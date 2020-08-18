// @Description
// @Author  renjingneng
// @CreateTime  2020/8/15 13:50
package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/renjingneng/a_simple_go_project/core/config"
	"github.com/renjingneng/a_simple_go_project/lib/utility"
	"github.com/renjingneng/a_simple_go_project/model"
	"strings"
)

// InitApp 初始化参数apptype、whitelist
//
// @Author  renjingneng
//
// @CreateTime  2020/8/15 16:23
func InitApp(ctx iris.Context) {
	path := ctx.Path()
	pathSlice := strings.Split(path, "/")
	apptype := pathSlice[1]
	whitelist := config.Total[utility.ToCamel(apptype)+"Whitelist"]
	ctx.Values().Set("apptype", apptype)
	ctx.Values().Set("whitelist", whitelist)
	ctx.Next()
}
func CheckSign(ctx iris.Context) {
	data := ctx.FormValues()
	newData := make(map[string]string)
	sign := ""
	whitelist, _ := ctx.Values().Get("whitelist").(map[string]string)
	for key, value := range data {
		if key == "sign" {
			sign = value[0]
			continue
		}
		newData[key] = value[0]
	}
	if newData["timestamp"] == "" || sign == "" || newData["appid"] == "" {
		res := &model.ResponseFail{Status: "fail", Reason: "参数不足-1！"}
		ctx.JSON(res)
		return
	}

	if whitelist[newData["appid"]] == "" {
		res := &model.ResponseFail{Status: "fail", Reason: "非法访问-1！"}
		ctx.JSON(res)
		return
	}
	appsecret := whitelist[newData["appid"]]
	newSign := utility.CreateSignWithUrlencode(newData, appsecret)
	if newSign == sign {
		ctx.Values().Set("appid", newData["appid"])
		ctx.Values().Set("appsecret", appsecret)
		ctx.Next()
	} else {
		res := &model.ResponseFail{Status: "fail", Reason: "非法访问-2！"}
		ctx.JSON(res)
		return
	}
}
