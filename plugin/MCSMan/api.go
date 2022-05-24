package MCSMan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	con "github.com/xrcuo/XAODBRGTR/confg"
)

var Azom = fmt.Sprintf(
	"http://%s:%d/api/protected_instance/stop?uuid=%s&remote_uuid=%s&apikey=%s",
	con.Ziod.Mcsman.Addr,
	con.Ziod.Mcsman.Port,
	con.Ziod.Mcsman.Uld,
	con.Ziod.Mcsman.Gld,
	con.Ziod.Mcsman.Apikey,
)

func init() {
	fmt.Println(con.Ziod.NickNames)
}

type MCsmw struct {
	Status string `json:"status,"`
	Data   struct {
		Host string `json:"instanceUuid"`
	} `json:"data"`
}

func MCSAF(url string) (C MC, err error) {
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	var M MCsmw
	fmt.Print(M)
	json.Unmarshal([]byte(body), &M)
	C.Status = M.Status
	C.Host = M.Data.Host
	return C, err
}

type MC struct {
	Status string
	Host   string
}
