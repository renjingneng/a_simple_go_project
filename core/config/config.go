package config

import (
	"flag"
	"github.com/renjingneng/a_simple_go_project/lib/log"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

var Config *config
var DatabaseMap map[string]string

type config struct {
	Env              string
	BaseURL          string `yaml:"BaseURL"`          // BaseURL
	Port             string `yaml:"Port"`             // 端口
	MysqlJiafuW      string `yaml:"MysqlJiafuW"`      // 嘉福mysql写数据库
	MysqlJiafuR      string `yaml:"MysqlJiafuR"`      // 嘉福mysql读数据库
	RedisJiafuSingle string `yaml:"RedisJiafuSingle"` // 嘉福redis单机实例
}

func LoadConfig() {
	var envFlag = flag.String("env", "dev", "请输入env参数,默认值为dev!")
	var filename string

	flag.Parse()
	Config = &config{}
	if *envFlag == "prod" {
		filename = "config/config-prod.yaml"
	} else if *envFlag == "dev" {
		filename = "config/config-dev.yaml"
	} else {
		filename = "config/config-dev.yaml"
	}
	if yamlFile, err := ioutil.ReadFile(filename); err != nil {
		log.Error(err)
	} else if err = yaml.Unmarshal(yamlFile, Config); err != nil {
		log.Error(err)
	}
	Config.Env = *envFlag
	DatabaseMap = make(map[string]string)
	DatabaseMap["MysqlJiafuW"] = Config.MysqlJiafuW
	DatabaseMap["MysqlJiafuR"] = Config.MysqlJiafuR
	DatabaseMap["RedisJiafuSingle"] = Config.RedisJiafuSingle
}
