package models

import (
	"github.com/astaxie/beego/config"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

type CalSend struct {
	Key  string `json:"key"`
	Date string `json:"date"`
}

type CalGet struct {
	ErrorCode int `json:"error_code"`
	Reason    string `json:"reason"`
	Result    map[string]CalResult `json:"result"`
}

type CalResult struct {
	Date        string
	LunarYear   string
	Lunar       string
	AnimalsYear string
	Suit        string
	Avoid       string
}

func (this *CalSend) SendAndGet() string {
	conf, conf_err := config.NewConfig("ini", "conf/juhe.conf")
	if conf_err != nil {
		panic(conf_err)
	}
	calconf, _ := conf.GetSection("calendar")
	//this.Date需要在创建的对象的时候手动赋值
	resp, _ := http.Get(calconf["address"] + "?key=" + calconf["key"] + "&date=" + this.Date)
	defer resp.Body.Close()
	body_bytes, _ := ioutil.ReadAll(resp.Body)
	cg := new(CalGet)
	json.Unmarshal(body_bytes, cg)
	result := cg.Result["data"]
	return "公历：" + result.Date + "\n农历：" + result.LunarYear + result.Lunar + "\n属相：" + result.AnimalsYear + "\n事宜：" + result.Suit + "\n禁忌：" + result.Avoid
}