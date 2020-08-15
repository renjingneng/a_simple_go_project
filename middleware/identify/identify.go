package identify

import (
	"github.com/kataras/iris/v12"
	"github.com/renjingneng/a_simple_go_project/lib/spew"
)

// CheckToken 检查是否有token
//
//ctx.Path() => 对应全路径 比如/identify/api/get/token
//ctx.FormValues() => 返回get和post参数,格式如下
//(map[string][]string) (len=3) {
//	(string) (len=3) "nnn": ([]string) (len=1 cap=1) {
//	(string) (len=4) "dfdf"
//},
//	(string) (len=5) "test1": ([]string) (len=1 cap=1) {
//	(string) (len=6) "dfdfdf"
//},
//	(string) (len=4) "dfdf": ([]string) (len=1 cap=1) {
//	(string) (len=5) "dfdfd"
//}
//}
//
// @Author  renjingneng
//
// @CreateTime  2020/8/15 15:21
func CheckToken(ctx iris.Context) {
	shareInformation := "this is a sharable information between handlers"
	data := ctx.FormValues()
	path := ctx.Path()
	ctx.Values().Set("info", shareInformation)
	spew.Dump(data)
	spew.Dump(path)
	ctx.Next()
}
func CheckLogin(ctx iris.Context) {

}
func InitUser(ctx iris.Context) {

}
