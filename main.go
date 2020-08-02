package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/renjingneng/a_simple_go_project/controller/admin/news"
	_ "github.com/renjingneng/a_simple_go_project/core"
	"github.com/renjingneng/a_simple_go_project/service"
)

func main() {
	app := iris.New()
	app.Get("/ping", pong)
	mvc.Configure(app.Party("/api"), func(m *mvc.Application) {
		// Register Dependencies.
		m.Register(
			service.NewGreetService,
		)
		m.Party("/news").Handle(new(news.ArticleController))
	})
	app.Listen(":8080")
}

func pong(ctx iris.Context) {
	ctx.WriteString("pong")
}
