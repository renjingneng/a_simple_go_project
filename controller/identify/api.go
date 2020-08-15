package identify

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/renjingneng/a_simple_go_project/middleware"
	"github.com/renjingneng/a_simple_go_project/model"
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
func (thisController *ApiController) TokenHandler() {
	res := &model.ResponseSucc{Status: "succ", Info: "成功"}
	thisController.Ctx.JSON(res)
}
