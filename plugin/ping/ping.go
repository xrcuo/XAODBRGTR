package ping

import (
	"fmt"

	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	con "github.com/xrcuo/XAODBRGTR/confg"
)

var pingx = "https://api.gmit.vip"

var (
	Pingx = "/Api/Ping?format=json&ip="
)

//Ping
func init() {
	zero.OnRegex(`^ping (.*?(.*))`).
		Handle(func(ctx *zero.Ctx) {
			ctx.Send(con.Ziod.Sxh)
			List := ctx.State["regex_matched"].([]string)[1]
			fmt.Print(pingx + Pingx + List + "&node=0")
			ZW, _ := ping(pingx + Pingx + List + "&node=0")
			if ZW.Ping_max == "" {
				ctx.Send(
					message.ReplyWithMessage(ctx.Event.MessageID,
						message.Text(ZW.Msg)))
			} else {
				ctx.Send(
					message.ReplyWithMessage(ctx.Event.MessageID,
						message.Text(ZW.Msg, "\n"),
						message.Text("请求域名: ", ZW.Host, "\n"),
						message.Text("PING节点: ", ZW.Node, "\n"),
						message.Text("IP地址: ", ZW.Ip, "\n"),
						message.Text("归属地：", ZW.Location, "\n"),
						message.Text("平均延迟: ", ZW.Ping_max, "\n"),
						message.Text(ZW.Domain_ip),
					))
			}
		})
}
