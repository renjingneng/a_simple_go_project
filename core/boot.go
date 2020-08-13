package core

import (
	iris "github.com/kataras/iris/v12"

	"github.com/renjingneng/a_simple_go_project/core/config"
	"github.com/renjingneng/a_simple_go_project/core/router"
)

func Boot() {
	//load config
	config.LoadConfig()
	//create iris instance
	app := iris.New()
	app.Get("/ping", pong)
	//attach mvc
	router.AttachMvc(app)
	//run app
	app.Listen(":8080")
}
func pong(ctx iris.Context) {
	ctx.WriteString("pong")
}
