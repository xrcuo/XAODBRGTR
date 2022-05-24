package con

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"

	yaml "gopkg.in/yaml.v2"
)

var Ziod Config

type Config struct {
	// 配置文件结构体
	Scex          string   `yaml:"opzqr"`
	SuperU        []int64  `yaml:"SuperUsers"`
	NickNames     []string `yaml:"NickNames"`
	CommandPrefix string   `yaml:"CommandPrefix"`
	Server        struct {
		Address string `yaml:"Address"`
		Port    int64  `yaml:"Port"`
		Token   string `yaml:"Token"`
	} `yaml:"Server"`
	Mcsman struct {
		Addr   string `yaml:"addr"`
		Port   int64  `yaml:"port"`
		Apikey string `yaml:"apikey"`
		Uld    string `yaml:"UID"`
		Gld    string `yaml:"GID"`
	} `yaml:"Mcsman"`
}

func init() {
	// 初始化配置文件
	configFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Error(err)
	}
	err = yaml.Unmarshal(configFile, &Ziod)
	if err != nil {
		log.Error(err)
	}
}
