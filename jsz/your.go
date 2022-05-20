package jsz

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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
