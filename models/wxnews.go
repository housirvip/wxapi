package models

import (
	"github.com/astaxie/beego/config"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strconv"
)

type wxNewsSend struct {
	Pno int `json:"pno"`
	Ps  int `json:"ps"`
}

type wxNewsGet struct {
	ErrorCode int `json:"error_code"`
	Reason    string `json:"reason"`
	Result    map[string][]wxNewsResult `json:"result"`
}

type wxNewsResult struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Source   string `json:"source"`
	FirstImg string `json:"firstImg"`
	Mark     string `json:"mark"`
	Url      string `json:"url"`
}

func (this *wxNewsSend) SendAndGet() string {
	conf, conf_err := config.NewConfig("ini", "conf/juhe.conf")
	if conf_err != nil {
		panic(conf_err)
	}
	wxNewsconf, _ := conf.GetSection("wxnews")
	if !(this.Ps > 1 && this.Ps < 50) {
		this.Ps = 5
	}
	if !(this.Pno > 1 && this.Pno < 50) {
		this.Pno = 1
	}
	resp, _ := http.Get(wxNewsconf["address"] + "?key=" + wxNewsconf["key"] + "&pno=" + strconv.Itoa(this.Pno) + "&ps=" + strconv.Itoa(this.Ps))
	defer resp.Body.Close()
	body_bytes, _ := ioutil.ReadAll(resp.Body)
	wng := new(wxNewsGet)
	json.Unmarshal(body_bytes, wng)
	list := wng.Result["list"]
	var result string
	for _, v := range list {
		result += v.Title + "\n来源：" + v.Source + "\n链接：" + v.Url + "\n"
	}
	return result
}
