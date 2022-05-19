package jsz

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//mc
func MCSAF(url string) (C MC, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return C, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return C, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return C, err
	}
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
