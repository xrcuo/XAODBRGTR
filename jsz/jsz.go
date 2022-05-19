package jsz

import (
	"encoding/base64"
)

func Base64Encoding(str string) string { //Base64编码
	src := []byte(str)
	res := base64.StdEncoding.EncodeToString(src) //将编码变成字符串
	return res
}

//mc
type MotdJavaInfo struct {
	Status    string  `json:"status"`    //服务器状态
	Host      string  `json:"host"`      //服务器Host
	Motd      string  `json:"motd"`      //Motd信息
	Agreement int     `json:"agreement"` //协议版本
	Version   string  `json:"version"`   //支持的游戏版本
	Online    float64 `json:"online"`    //在线人数
	Max       float64 `json:"max"`       //最大在线人数
	Sample    []struct {
		Id   string `json:"id"`   //在线玩家ID
		Name string `json:"name"` //在线玩家名字
	} `json:"sample"` //在线玩家列表
	Favicon string `json:"favicon"` //服务器图标
	Delay   int64  `json:"delay"`   //连接延迟
}

type AWSQ struct {
	ImageUrl string `json:"imageUrl"`
}

//Ping
type PINGZAQ struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Ip        string `json:"ip"`
		Node      string `json:"node"`
		Host      string `json:"host"`
		Domain_ip string `json:"domain_ip"`
		Ping_min  string `json:"ping_min"`
		Ping_avg  string `json:"ping_avg"`
		Ping_max  string `json:"Ping_max"`
		Location  string `json:"location"`
	} `json:"data"`
}

type SXRQZ struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Domain      string `json:"domain"`
		ServiceName string `json:"serviceName"`
		HomeUrl     string `json:"homeUrl"`
		Class       string `json:"class"`
		UnitName    string `json:"unitName"`
		Icp         string `json:"icp"`
		Time        string `json:"time"`
	} `json:"data"`
}

type MotdBEInfo struct {
	Status    string  `json:"status"`     //服务器状态
	Host      string  `json:"host"`       //服务器Host
	Motd      string  `json:"motd"`       //Motd信息
	Agreement int     `json:"agreement"`  //协议版本
	Version   string  `json:"version"`    //支持的游戏版本
	Online    float64 `json:"online"`     //在线人数
	Max       float64 `json:"max"`        //最大在线人数
	LevelName string  `json:"level_name"` //存档名字
	GameMode  string  `json:"gamemode"`   //游戏模式
	Delay     int64   `json:"delay"`      //连接延迟
}

type Edce struct {
	Success string `json:"success"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Xuid string `json:"xuid"`
	} `json:"Data"`
}

//天气
type XUJXWSAA struct {
	Code        string `json:"code"`
	Msg         string `json:"msg"`
	City        string `json:"city"`
	Update_time string `json:"update_time"`
	Wea         string `json:"wea"`
	Wea_img     string `json:"wea_img"`
	Tem         string `json:"tem"`
	Tem_day     string `json:"tem_day"`
	Tem_night   string `json:"tem_night"`
	Win         string `json:"win"`
	Win_speed   string `json:"win_speed"`
	Win_meter   string `json:"win_meter"`
	Air         string `json:"air"`
	Time        string `json:"time"`
}

//疫情
type ZXs_xdzss struct {
	Code string `json:"code"`
	Text string `json:"text"`
	Data struct {
		Time      string `json:"time"`
		Name      string `json:"name"`
		Suspected string `json:"suspected"`
		Death     string `json:"death"`
		Cure      string `json:"cure"`
		Totaladd  string `json:"totaladd"`
		Econ      string `json:"econ"`
		Asymptom  string `json:"asymptom"`
	} `json:"data"`
}

type AobileInfo struct {
	Success  bool   `json:"success"`
	Status   int    `json:"status"`
	Message  string `json:"message"`
	Version  string `json:"version"`
	Codename string `json:"codename"`
	Data     struct {
		Exist bool `json:"exist"`
		Info  []struct {
			Uuid     string   `json:"uuid"`
			Info     string   `json:"info"`
			Nuid     string   `json:"nuid"`
			Black_id string   `json:"black_id"`
			Qq       int      `json:"qq"`
			Name     string   `json:"name"`
			Level    int      `json:"level"`
			Photos   []string `json:"photos"`
		} `json:"info"`
	} `json:"data"`
}
