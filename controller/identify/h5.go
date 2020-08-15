package identify

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type H5Controller struct {
	Ctx iris.Context
}

func (thisController *H5Controller) BeforeActivation(mvc mvc.BeforeActivation) {

}
func (thisController *H5Controller) AfterActivation(mvc mvc.AfterActivation) {

}
func (thisController *H5Controller) BeginRequest(ctx iris.Context) {

}
func (thisController *H5Controller) EndRequest(ctx iris.Context) {

}

/*func (thisController *H5Controller) MyCustomHandler() {

}
func (thisController *H5Controller) PostLogin() {

}*/
