package news

import (
	"github.com/renjingneng/a_simple_go_project/data/mysql"
	"github.com/renjingneng/a_simple_go_project/data/redis"
)

type Article struct {
	country          string
	db               *mysql.Jiafu
	localRedisClient *redis.Jiafu

	*Base
}

func NewArticle(country string, name string) *Article {
	a := &Article{
		country:          country,
		db:               mysql.NewJiafu(),
		Base:             NewBase(name),
		localRedisClient: redis.NewJiafu(),
	}
	return a
}
func (a *Article) FetchAuthorById(id string) map[string]interface{} {
	a.db.SetTablename("author")
	return a.db.FetchRow("first_name,last_name", map[string]string{"id": id})
}
func (a *Article) FetchLocalCache() interface{} {
	if res, err := a.localRedisClient.Get("test123"); err != nil {
		panic(err)
	} else {
		return res
	}
}

func (a *Article) GetTitleByName(name string) string {
	return "title is :" + name
}
