package core

import (
	"flag"
	ioutil "io/ioutil"

	log "github.com/renjingneng/a_simple_go_project/lib/log"
	yaml "gopkg.in/yaml.v2"
)

//Config dfdf
var Config *ConfigT

//ConfigT dfdf
type ConfigT struct {
	Env      string
	BaseUrl  string `yaml:"BaseUrl"`  // base url
	Port     string `yaml:"Port"`     // 端口
	LogFile  string `yaml:"LogFile"`  // 日志文件
	MySqlUrl string `yaml:"MySqlUrl"` // 数据库连接地址
}

//LoadConfig dfdf
func LoadConfig() {
	var envFlag = flag.String("env", "normal", "请输入env参数!")
	flag.Parse()
	Config = &ConfigT{}
	var filename string
	if *envFlag == "prod" {
		filename = "config-prod.yaml"
	} else if *envFlag == "dev" {
		filename = "config-dev.yaml"
	} else {
		filename = "config.yaml"
	}
	if yamlFile, err := ioutil.ReadFile(filename); err != nil {
		log.Error(err)
	} else if err = yaml.Unmarshal(yamlFile, Config); err != nil {
		log.Error(err)
	}
	Config.Env = *envFlag

}
