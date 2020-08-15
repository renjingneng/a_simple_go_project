package identify

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/renjingneng/a_simple_go_project/middleware"
)

type ApiController struct {
	Ctx iris.Context
}

func (thisController *ApiController) BeforeActivation(mvc mvc.BeforeActivation) {
	mvc.Handle(
		"POST",
		"/token",
		"TokenHandler",
		middleware.CheckSign,
	)
}
func (thisController *ApiController) AfterActivation(mvc mvc.AfterActivation) {

}
func (thisController *ApiController) BeginRequest(ctx iris.Context) {

}
func (thisController *ApiController) EndRequest(ctx iris.Context) {

}
func (thisController *ApiController) TokenHandler() {

}
func (thisController *ApiController) PostLogin() {

}
