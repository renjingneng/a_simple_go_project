package core

import (
	iris "github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	"github.com/renjingneng/a_simple_go_project/controller/admin/news"
	"github.com/renjingneng/a_simple_go_project/core/config"
	"github.com/renjingneng/a_simple_go_project/service"
)

func Boot() {
	//load config
	config.LoadConfig()
	//create iris instance
	app := iris.New()
	app.Get("/ping", pong)
	//attach mvc
	attachMvc(app)
	//run app
	app.Listen(":8080")
}
func pong(ctx iris.Context) {
	ctx.WriteString("pong")
}
func attachMvc(app *iris.Application) {
	mvc.Configure(app.Party("/api"), func(m *mvc.Application) {
		// Register Dependencies on api
		m.Register(
			service.NewGreetService,
		)
		mvcNews := m.Party("/news")
		// Register Dependencies on news
		/*mvcNews.Register(
			service.NewGreetService,
		)*/
		mvcNews.Party("/article").Handle(new(news.ArticleController))
	})
}
