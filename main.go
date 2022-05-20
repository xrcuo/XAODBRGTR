package main

import (
	"fmt"
	"io/ioutil"

	//-----------------------------------------------//
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/driver"
	"gopkg.in/yaml.v2"

	//-----------------------------------------------//
	_ "github.com/xrcuo/XAODBRGTR/plugin/mc"
)

var config Config

func init() {
	log.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[zero][%time%][%lvl%]: %msg% \n",
	})

	// 初始化配置文件
	configFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Error(err)
	}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Error(err)
	}
}

func main() {
	zero.OnCommand("hello").
		Handle(func(ctx *zero.Ctx) {
			ctx.Send("world")
		})

	zero.RunAndBlock(zero.Config{
		NickName:      []string{"bot"},
		CommandPrefix: "/",
		SuperUsers:    []int64{123456},
		Driver: []zero.Driver{
			driver.NewWebSocketClient(
				fmt.Sprintf("ws://%s:%d", config.Server.Address, config.Server.Port),
				config.Server.Token,
			)},
	}, nil)
}
