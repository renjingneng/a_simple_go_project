package identify

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type ApiController struct {
	Ctx iris.Context
}

func (thisController *ApiController) BeforeActivation(mvc mvc.BeforeActivation) {
}
func (thisController *ApiController) GetGetToken() {

}
