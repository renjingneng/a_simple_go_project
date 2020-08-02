package news

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	model "github.com/renjingneng/a_simple_go_project/model/news"
	"github.com/renjingneng/a_simple_go_project/service"
	servicenews "github.com/renjingneng/a_simple_go_project/service/news"
)

// ArticleController handles the index.
type ArticleController struct {
	ServiceGreet       service.Greet
	ServiceNewsArticle *servicenews.Article
	Ctx                iris.Context
}

func (a *ArticleController) BeforeActivation(b mvc.BeforeActivation) {
	a.ServiceNewsArticle = servicenews.NewArticle("china", "renjingneng")
}

// Get serves [GET] /.
// Query: name
func (a *ArticleController) Get(req model.Request) (model.Response, error) {
	message, err := a.ServiceGreet.Say(req.Name)
	if err != nil {
		return model.Response{}, err
	}
	message = message + a.ServiceNewsArticle.GetTitleByName(req.Name)
	return model.Response{Message: message}, nil
}
