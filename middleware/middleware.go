// @Description
// @Author  renjingneng
// @CreateTime  2020/8/15 13:50
package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/renjingneng/a_simple_go_project/lib/spew"
)

func CheckSign(ctx iris.Context) {
	data := ctx.FormValues()
	path := ctx.Path()
	spew.Dump(data)
	spew.Dump(path)
	ctx.Next()
}
