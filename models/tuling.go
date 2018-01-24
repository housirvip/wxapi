package models

import (
	"github.com/astaxie/beego/config"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"strconv"
	"fmt"
)

type TulingSend struct {
	Key    string `json:"key"`
	Info   string `json:"info"`
	Userid string `json:"userid"`
}

type TulingGet struct {
	ErrorCode int `json:"error_code"`
	Reason    string `json:"reason"`
	Result    TulingResult `json:"result"`
}

type TulingResult struct {
	Code int `json:"code"`
	Text string `json:"text"`
}

func (this *TulingSend) SendAndGet() string {
	conf, conf_err := config.NewConfig("ini", "conf/juhe.conf")
	if conf_err != nil {
		panic(conf_err)
	}
	tuling, _ := conf.GetSection("tuling")
	resp, _ := http.Get(tuling["address"] + "?info=" + this.Info + "&userid=" + this.Userid + "&key=" + tuling["key"])
	defer resp.Body.Close()
	body_bytes, _ := ioutil.ReadAll(resp.Body)
	tg := new(TulingGet)
	json.Unmarshal(body_bytes, tg)
	if tg.ErrorCode != 0 {
		log := Log{Content:strconv.Itoa(tg.ErrorCode) + "聚合图灵错误:" + tg.Reason}
		log.AddLog()
		fmt.Println(tg)
		return "兄弟，你确定不是在跟我玩蛇？"
	}
	return tg.Result.Text
}






