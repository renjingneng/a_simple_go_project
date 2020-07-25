package main

import (
	"fmt"

	core "github.com/renjingneng/a_simple_go_project/core"
)

func main() {
	core.LoadConfig()
	/****test start***/
	fmt.Println(core.Config.Env)
	/****test end***/
}
