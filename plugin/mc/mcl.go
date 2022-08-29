package mc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

//mc
func MCSAF(url string) (C MC, err error) {
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	var M MotdBEInfo
	fmt.Print(M)
	json.Unmarshal([]byte(body), &M)
	C.Status = M.Status
	C.Host = M.Host
	C.Motd = M.Motd
	C.Agreement = M.Agreement
	C.Version = M.Version
	C.Online = M.Online
	C.Max = M.Max
	C.LevelName = M.LevelName
	C.GameMode = M.GameMode
	C.Delay = M.Delay
	return C, err
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
