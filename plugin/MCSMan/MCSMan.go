package MCSMan

import (
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	con "github.com/xrcuo/XAODBRGTR/confg"
)

func init() {
	zero.OnFullMatchGroup([]string{"检查身体", "自检", "启动自检", "系统状态"}, zero.AdminPermission).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.Send("少女祈祷中......\n查什么查，别查了...")
			MC, _ := MCSAF(Azom)
			if MC.Host != con.Ziod.Mcsman.Uld {
				ctx.Send("查询失败\n请检查IP地址跟端口")
			} else {
				ctx.Send(
					message.ReplyWithMessage(ctx.Event.MessageID,
						message.Text("请求成功，正在关闭服务器...\n", MC.Host),
					),
				)
			}
		})
}
