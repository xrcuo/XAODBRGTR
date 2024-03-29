package mc

import (
	"fmt"

	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	con "github.com/xrcuo/XAODBRGTR/confg"
)

var (
	xk = "/api?host="
	zw = "/api/java?host="
)

var (
	mcst = "/status_img?host="
	mcsa = "/status_img/java?host="
)

func init() {
	zero.OnRegex(`^mc (.*?(.*))`).
		Handle(func(ctx *zero.Ctx) {
			ctx.Send(con.Ziod.Sxh)
			List := ctx.State["regex_matched"].([]string)[1]
			fmt.Print(con.Ziod.Stop + xk + List)
			MC, _ := MCSAF(con.Ziod.Stop + xk + List)
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
						message.Image(con.Ziod.Stop+mcst+List),
						message.Text("服务器状态：", MC.Status, "\n"),
						message.Text("协议版本：", MC.Agreement, "\n"),
						message.Text("游戏版本：", MC.Version, "\n"),
						message.Text("描述文本：", MC.Motd, "\n"),
						message.Text("在线人数：", MC.Online, "/", MC.Max),
						message.Text("(", c, "%", ")", "\n"),
						message.Text("游戏模式：", MC.GameMode, "\n"),
						message.Text("游戏延迟：", MC.Delay, ".", s, "ms"),
					),
				)
			}
		})
}

func init() {
	zero.OnRegex(`^mcl (.*?(.*))`).
		Handle(func(ctx *zero.Ctx) {
			ctx.Send(con.Ziod.Sxh)
			List := ctx.State["regex_matched"].([]string)[1]
			fmt.Print(con.Ziod.Stop + zw + List)
			MC, _ := MCSAF(con.Ziod.Stop + zw + List)
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
						message.Image(con.Ziod.Stop+mcsa+List),
						message.Text("服务器状态：", MC.Version, "\n"),
						message.Text("协议版本：", MC.Motd, "\n"),
						message.Text("游戏版本：", MC.Agreement, "\n"),
						message.Text("描述文本：", MC.Status, "\n"),
						message.Text("在线人数：", MC.Online, "/", MC.Max),
						message.Text("(", c, "%", ")", "\n"),
						message.Text("游戏延迟：", MC.Delay, ".", s, "ms"),
					),
				)
			}
		})
}
