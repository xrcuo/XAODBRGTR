package main

import (
	"fmt"

	//-----------------------------------------------//
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/driver"
	con "github.com/xrcuo/XAODBRGTR/confg"

	//-----------------------------------------------//
	//_ "github.com/xrcuo/XAODBRGTR/plugin/MCSMan"
	_ "github.com/xrcuo/XAODBRGTR/plugin/mc"
)

func init() {
	log.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[zero][%time%][%lvl%]: %msg% \n",
	})

}

func main() {

	zero.OnCommand("hello").
		Handle(func(ctx *zero.Ctx) {
			ctx.Send("world")
		})

	zero.RunAndBlock(zero.Config{
		NickName:      con.Ziod.Wsn.NickNames,
		CommandPrefix: con.Ziod.Wsn.CommandPrefix,
		SuperUsers:    con.Ziod.Wsn.SuperU,
		Driver: []zero.Driver{
			driver.NewWebSocketClient(
				fmt.Sprintf("ws://%s:%d", con.Ziod.Server.Address, con.Ziod.Server.Port),
				con.Ziod.Server.Token,
			)},
	}, nil)

}
