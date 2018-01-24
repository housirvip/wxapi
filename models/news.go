package models

import (
	"github.com/astaxie/beego/config"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type NewsSend struct {
	Type string `json:"type"`
}

type NewsGet struct {
	ErrorCode int `json:"error_code"`
	Reason    string `json:"reason"`
	Result    map[string][]NewsResult `json:"result"`
}

type NewsResult struct {
	Title      string `json:"title"`
	Date       string `json:"date"`
	AuthorName string `json:"author_name"`
	Realtype   string `json:"realtype"`
	Url        string `json:"url"`
}

func (this *NewsSend) SendAndGet() string {
	conf, conf_err := config.NewConfig("ini", "conf/juhe.conf")
	if conf_err != nil {
		panic(conf_err)
	}
	newsconf, _ := conf.GetSection("news")
	resp, _ := http.Get(newsconf["address"] + "?key=" + newsconf["key"] + "&type=" + this.Type)
	defer resp.Body.Close()
	body_bytes, _ := ioutil.ReadAll(resp.Body)
	ng := new(NewsGet)
	json.Unmarshal(body_bytes, ng)
	data := ng.Result["data"]
	var result string
	for k, v := range data {
		result += v.Title + "\n作者：" + v.AuthorName + "\n时间：" + v.Date + "\n类型：" + v.Realtype + "\n链接：" + v.Url + "\n"
		if k > 4 {
			break
		}
	}
	return result
}
