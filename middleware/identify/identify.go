package identify

import "github.com/kataras/iris/v12"

func CheckToken(ctx iris.Context) {
	shareInformation := "this is a sharable information between handlers"

	requestPath := ctx.Path()
	println("Before the mainHandler: " + requestPath)

	ctx.Values().Set("info", shareInformation)
	ctx.Next() // execute the next handler, in this case the main one.
}
func CheckLogin(ctx iris.Context) {

}
func InitUser(ctx iris.Context) {

}
