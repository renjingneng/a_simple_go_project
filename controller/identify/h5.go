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
func (thisController *H5Controller) GetLogin() {
	//info := thisController.Ctx.Values().GetString("info")
}
func (thisController *H5Controller) GetIdentifyCard() {

}
func (thisController *H5Controller) GetCardFinished() {

}
func (thisController *H5Controller) GetIdentifyFace() {

}
func (thisController *H5Controller) GetIdentifyFinished() {

}
func (thisController *H5Controller) GetInfo1() {

}
func (thisController *H5Controller) GetInfo2() {

}
