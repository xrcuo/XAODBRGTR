package mc

import (
	"fmt"

	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"github.com/xrcuo/XAODBRGTR/api"
	"github.com/xrcuo/XAODBRGTR/jsz"
)

func init() {
	zero.OnRegex(`^mco (.*?(.*))`).
		Handle(func(ctx *zero.Ctx) {
			ctx.Send("少女祈祷中......\n查什么查，别查了...")
			List := ctx.State["regex_matched"].([]string)[1]
			fmt.Print(api.APIC + api.MCPE + List)
			MC, _ := jsz.MCSAF(api.APIC + api.MCPE + List)
			if MC.Host == "" {
				ctx.Send("查询失败\n请检查IP地址跟端口")
			} else {
				a := (float64(MC.Online) * 32)
				b := (float64(MC.Max) * 32)
				r := (MC.Delay)
				c := (float32(a/b) * 100)
				s := (float64(r) * 15)
				ctx.Send(
					message.ReplyWithMessage(ctx.Event.MessageID,
						message.Text("[MCBE服务器信息]\n"),
						message.Text("服务器状态：", MC.Status, "\n"),
						message.Text("协议版本：", MC.Motd, "\n"),
						message.Text("游戏版本：", MC.Agreement, "\n"),
						message.Text("描述文本：", MC.Version, "\n"),
						message.Text("在线人数：", MC.Online, "/", MC.Max),
						message.Text("(", c, "%", ")", "\n"),
						message.Text("游戏模式：", MC.GameMode, "\n"),
						message.Text("游戏延迟：", MC.Delay, ".", s, "ms"),
					),
				)
			}
		})
}
