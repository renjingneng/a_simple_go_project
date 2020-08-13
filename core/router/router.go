package router

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/renjingneng/a_simple_go_project/controller/admin/news"
	"github.com/renjingneng/a_simple_go_project/controller/identify"
	"github.com/renjingneng/a_simple_go_project/service"
)

func AttachMvc(app *iris.Application) {
	mvc.Configure(app.Party("/admin"), func(m *mvc.Application) {
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
	mvc.Configure(app.Party("/identify"), func(m *mvc.Application) {
		m.Party("/api").Handle(new(identify.ApiController))
		m.Party("/h5").Handle(new(identify.H5Controller))
	})
}
