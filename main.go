package main

import (
	"github.com/renjingneng/a_simple_go_project/core"
)

// main 应用初始化
//
// @Author  renjingneng
//
// @CreateTime  2020/8/13 16:19
func main() {
	core.Boot()
}

/*// test 测试用的
//
// @Author  renjingneng
//
// @CreateTime  2020/8/18 9:46
func test() {
	config.LoadConfig()
	service := service.NewIdentify()
	info,_:=service.GetInfoByToken("3bcda3cfd1f3b0f6d881dc2ec88d70ed")
	spew.Dump(info)
}*/
