package news

import (
	"github.com/renjingneng/a_simple_go_project/data/mysql"
	"github.com/renjingneng/a_simple_go_project/data/redis"
)

type Article struct {
	country          string
	db               *mysql.Jiafu
	localRedisClient *redis.Base

	*base
}

func NewArticle(country string, name string) *Article {
	a := &Article{
		country:          country,
		db:               mysql.NewJiafu(),
		base:             NewBase(name),
		localRedisClient: redis.NewBase("LocalRedis", "Single"),
	}

	return a
}
func (a *Article) FetchLatest() interface{} {
	a.db.SetTablename("author")
	return a.db.FetchRow("first_name,last_name", map[string]string{"id": "27"})
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
