package main

import (
	"fmt"
	"github.com/renjingneng/a_simple_go_project/model/news"
)

func main() {
	article := news.NewArticle("china", "ren")
	/****test start***/
	fmt.Println(article.FetchLocalCache())
	/****test end***/
}
