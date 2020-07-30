package core

import (
	"flag"
	"io/ioutil"

	log "github.com/renjingneng/a_simple_go_project/lib/log"
	yaml "gopkg.in/yaml.v2"
)

//Config is
var Config *config

//DatabaseMap is
var DatabaseMap map[string]string

type config struct {
	Env              string
	BaseURL          string `yaml:"BaseURL"`          // BaseURL
	Port             string `yaml:"Port"`             // 端口
	LocalJiafuW      string `yaml:"LocalJiafuW"`      // 数据库连接地址
	LocalJiafuR      string `yaml:"LocalJiafuR"`      // 数据库连接地址
	LocalRedisSingle string `yaml:"LocalRedisSingle"` // 缓存地址
}

func loadConfig() {
	var envFlag = flag.String("env", "normal", "请输入env参数!")
	flag.Parse()
	Config = &config{}
	var filename string
	if *envFlag == "prod" {
		filename = "config/config-prod.yaml"
	} else if *envFlag == "dev" {
		filename = "config/config-dev.yaml"
	} else {
		filename = "config/config.yaml"
	}
	if yamlFile, err := ioutil.ReadFile(filename); err != nil {
		log.Error(err)
	} else if err = yaml.Unmarshal(yamlFile, Config); err != nil {
		log.Error(err)
	}
	Config.Env = *envFlag
	DatabaseMap = make(map[string]string)
	DatabaseMap["LocalJiafuW"] = Config.LocalJiafuW
	DatabaseMap["LocalJiafuR"] = Config.LocalJiafuR
	DatabaseMap["LocalRedisSingle"] = Config.LocalRedisSingle
}
func init() {
	loadConfig()
}
