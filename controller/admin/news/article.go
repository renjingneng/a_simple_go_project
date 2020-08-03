package news

import (
	"fmt"
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

func (a *ArticleController) Get(req model.Request) (model.Response, error) {
	message, err := a.ServiceGreet.Say(req.Name)
	if err != nil {
		return model.Response{}, err
	}
	message = message + a.ServiceNewsArticle.GetTitleByName(req.Name)
	return model.Response{Message: message}, nil
}
func (a *ArticleController) GetAuthor(req model.Author) (model.Response, error) {
	author := a.ServiceNewsArticle.FetchAuthorById(req.Id)
	first_name := fmt.Sprintf("%v", author["first_name"])
	last_name := fmt.Sprintf("%v", author["last_name"])
	return model.Response{Message: "My firstname is" + first_name + ",My lastname is " + last_name}, nil
}
