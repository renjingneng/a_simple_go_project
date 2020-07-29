package news

import (
	"github.com/renjingneng/a_simple_go_project/data/mysql"
)

type article struct {
	country string
	db      *mysql.Jiafu

	*base
}

func NewArticle(country string, name string) *article {
	a := &article{
		country: country,
		db:      mysql.NewJiafu(),
		base:    NewBase(name),
	}
	return a
}
func (a *article) FetchLatest() interface{} {
	a.db.SetTablename("author")
	return a.db.FetchRow("first_name,last_name", map[string]string{"id": "27"})
}
