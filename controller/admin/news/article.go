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
	firstName, ok := author["first_name"].(string)
	if ok != true {
		return model.Response{Message: "first_name 断言失败"}, nil
	}
	lastName, ok := author["last_name"].(string)
	if ok != true {
		return model.Response{Message: "last_name 断言失败"}, nil
	}
	/*firstName := fmt.Sprintf("%v", author["first_name"])
	lastName := fmt.Sprintf("%v", author["last_name"])*/
	return model.Response{Message: "My firstname is" + firstName + ",My lastname is " + lastName}, nil
}
