package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	yaml "gopkg.in/yaml.v2"

	"github.com/renjingneng/a_simple_go_project/lib/log"
)

//全局配置变量
var Config map[string]string
var DatabaseMap map[string]string
var Hostlist map[string]string
var SensetimeWhitelist map[string]string

// LoadConfig 载入配置,第一版配置变量Config等用struct类型，不太灵活，第二版改成了map类型，比较灵活
//
// @Author  renjingneng
//
// @CreateTime  2020/8/13 18:42
func LoadConfig() {
	//初始化变量
	var envFlag = flag.String("env", "local", "请输入env参数,默认值为local!")
	flag.Parse()
	var filename string
	Config = make(map[string]string)
	DatabaseMap = make(map[string]string)
	SensetimeWhitelist = make(map[string]string)
	Hostlist = make(map[string]string)
	yamlMap := make(map[interface{}]interface{})
	if *envFlag == "prod" {
		filename = "config/config-prod.yaml"
	} else if *envFlag == "dev" {
		filename = "config/config-dev.yaml"
	} else {
		*envFlag = "local"
		filename = "config/config-local.yaml"
	}
	if yamlFile, err := ioutil.ReadFile(filename); err != nil {
		log.Error(err)
	} else if err = yaml.Unmarshal(yamlFile, &yamlMap); err != nil {
		log.Error(err)
	}

	//赋值配置变量
	Config["Env"] = *envFlag
	for key, value := range yamlMap {
		keyStr := fmt.Sprint(key)
		if strings.HasPrefix(keyStr, "Mysql") || strings.HasPrefix(keyStr, "Redis") {
			DatabaseMap[keyStr] = fmt.Sprint(value)
			continue
		}
		if keyStr == "SensetimeWhitelist" {
			valueStr, _ := yamlMap[keyStr].(map[interface{}]interface{})
			//valueStr,_:=m[keyStr].(map[string]string)  不能这么用，这样直接转换会有问题
			for key, value := range valueStr {
				key := fmt.Sprint(key)
				value := fmt.Sprint(value)
				SensetimeWhitelist[key] = value
			}
			continue
		}
		if keyStr == "Hostlist" {
			valueStr, _ := yamlMap[keyStr].(map[interface{}]interface{})
			//valueStr,_:=m[keyStr].(map[string]string)  不能这么用，这样直接转换会有问题
			for key, value := range valueStr {
				key := fmt.Sprint(key)
				value := fmt.Sprint(value)
				Hostlist[key] = value
			}
			continue
		}
		Config[keyStr] = fmt.Sprint(value)
	}
}
