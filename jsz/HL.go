package jsz

type MC struct {
	Status    string  //服务器状态
	Host      string  //服务器Host
	Motd      string  //Motd信息
	Agreement int     //协议版本
	Version   string  //支持的游戏版本
	Online    float64 //在线人数
	Max       float64 //最大在线人数
	LevelName string  //存档名字
	GameMode  string  //游戏模式
	Delay     int64
}
