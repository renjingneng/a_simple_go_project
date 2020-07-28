package main

import (
	"fmt"

	core "github.com/renjingneng/a_simple_go_project/core"
	container "github.com/renjingneng/a_simple_go_project/core/container"
	mysql "github.com/renjingneng/a_simple_go_project/data/mysql"
)

func main() {
	dbptr := container.GetEntityFromMysqlContainer("LocalJiafu", "W")
	test := mysql.base{
		Tablename: "author",
		Dbname:    "jiafu",
		Dbptr:     dbptr,
	}
	/****test start***/
	fmt.Println(core.Config.LocalJiafuR)
	fmt.Println(dbptr)
	fmt.Println(test.FetchRow("first_name,last_name", map[string]string{"id": "27"}))
	/****test end***/
}
