package identify

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/jinzhu/now"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	"github.com/renjingneng/a_simple_go_project/core/config"
	"github.com/renjingneng/a_simple_go_project/middleware"
	"github.com/renjingneng/a_simple_go_project/model"
	service "github.com/renjingneng/a_simple_go_project/service/identify"
)

type ApiController struct {
	Ctx     iris.Context
	salt    string
	service *service.Identify
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
	thisController.salt = "dfwer"
	thisController.service = service.NewIdentify()
}
func (thisController *ApiController) EndRequest(ctx iris.Context) {

}
func (thisController *ApiController) TokenHandler() {
	urlCheck := thisController.Ctx.FormValue("url_check")
	urlResult := thisController.Ctx.FormValue("url_result")
	appid := thisController.Ctx.FormValue("appid")
	if urlCheck == "" || urlResult == "" {
		res := &model.ResponseFail{Status: "fail", Reason: "参数不足！"}
		thisController.Ctx.JSON(res)
		return
	}
	BeginningOfDay := fmt.Sprint(now.BeginningOfDay().Unix())
	ttl := now.EndOfDay().Unix() - time.Now().Unix()
	fullStr := urlCheck + urlResult + thisController.salt + appid + BeginningOfDay
	token := fmt.Sprintf("%x", md5.Sum([]byte(fullStr)))
	info, _ := thisController.service.GetInfoByToken(token)
	if info == nil || len(info) == 0 {
		info = map[string]string{
			"urlCheck":  urlCheck,
			"urlResult": urlResult,
			"appid":     appid,
		}
		thisController.service.StoreInfoByToken(token, info, time.Duration(ttl)*time.Second)
	}
	//成功之后返回网址
	url := config.Config["Baseurl"] + "/identify/h5/login?token+" + token
	res := &model.ResponseSucc{Status: "succ", Info: url}
	thisController.Ctx.JSON(res)
}
