package models

import (
	"encoding/xml"
	"strings"
	"github.com/astaxie/beego"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"
)

const (
	EVENT = "event"
	TEXT = "text"
	BAN = "ban"
	NULL = "null"
	MENU = "menu"
	SUBSCRIBE = "subscribe"
	UNSUBSCRIBE = "unsubscribe"
	IMAGE = "image"
	OK = "ok"
	MAN = 1
	WOMAN = 0
)

type Wxmsg struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string  `xml:"ToUserName"`
	FromUserName string  `xml:"FromUserName"`
	CreateTime   string  `xml:"CreateTime"`
	MsgType      string  `xml:"MsgType"`
	Content      string  `xml:"Content"`
	MsgId        string  `xml:"MsgId"`
	Event        string `xml:"Event"`
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int `json:"expires_in"`
}

type UserInfo struct {
	Errcode    int `json:"errcode"`
	Errmsg     string `json:"errmsg"`
	//UserInfoList []interface{} `json:"user_info_list"`
	Subscribe  int `json:"subscribe"`
	Openid     string `json:"openid"`
	Nickname   string `json:"nickname"`
	Sex        int `json:"sex"`
	Language   string `json:"language"`
	City       string `json:"city"`
	Province   string `json:"province"`
	Country    string `json:"country"`
	Headimgurl string `json:"headimgurl"`
}

func (this *Wxmsg) Xml2Msg(xml_bytes []byte) {
	err := xml.Unmarshal(xml_bytes, this)
	if err != nil {
		panic(err)
	}
}

func (this *Wxmsg) Str2Msg() string {
	this.ToUserName, this.FromUserName = this.FromUserName, this.ToUserName
	result, err := xml.MarshalIndent(this, "", " ")
	if err != nil {
		panic(err)
	}
	return string(result)
}

func (this *Wxmsg) Match() {
	switch this.MsgType {
	case TEXT:
		//fmt.Println(this.Content)
		switch  {
		case ifBan(this.Content):
			this.Content = "对不起，您输入的字符含有禁忌字符"
		case ifLol(this.Content):
			to_search := strings.Split(this.Content, "#")
			to_search_len := len(to_search)
			lolinfo := new(LolInfo)
			lolinfo.Gameid = to_search[to_search_len - 1]
			lolinfo.Server = to_search[to_search_len - 2]
			this.Content = lolinfo.GetInfo()
			//fmt.Println(lolinfo)
		case ifMenu(this.Content):
			this.Content="您可以回复“天气”“黄历”“新闻”“微信精选”“菜谱”“笑话”“段子”……\n" +
				"或者“lol#游戏大区#游戏ID”查询隐藏分"
		case ifCalendar(this.Content):
			cs:=new(CalSend)
			to_search := strings.Split(this.Content, "#")
			if len(to_search)==1 {
				cs.Date=time.Now().Format("2006-01-02")
				this.Content="(示例:黄历#"+cs.Date+")\n"+cs.SendAndGet()
			}else {
				cs.Date=to_search[1]
				this.Content=cs.SendAndGet()
			}
		case ifwxNews(this.Content):
			wns:=new(wxNewsSend)
			this.Content=wns.SendAndGet()
		case ifNews(this.Content):
			ns:=new(NewsSend)
			this.Content=ns.SendAndGet()
		default:
			ts := new(TulingSend)
			ts.Info = this.Content
			ts.Userid = strings.Replace(strings.Replace(this.FromUserName, "_", "0", -1), "-", "0", -1)
			this.Content = ts.SendAndGet()
		}

	case EVENT:
		user := new(User)
		user.Uid = this.FromUserName
		log := new(Log)
		switch this.Event {
		case SUBSCRIBE:
			this.MsgType = TEXT
			this.Content = "感谢关注" + beego.AppConfig.String("sitename") +
				"\n查看菜单请发送‘菜单’或者？"
			user_info := getUser(this.FromUserName)
			//fmt.Println(user_info,user_info.Errcode)
			if user_info.Errcode == 0 {
				user.Name = user_info.Nickname
				user.City = user_info.City
				user.Province = user_info.Province
				user.Country = user_info.Country
				user.Sex = user_info.Sex
				user.Language = user_info.Language
			} else {
				log.Content = "微信access_token错误" + user_info.Errmsg
			}
			user.AddUser()
			log.Content += "新增一个关注者" + user.Uid
		case UNSUBSCRIBE:
			user.DelUser()
			log.Content = "减少一个关注者" + user.Uid
		}
		log.AddLog()
	default:
		this.MsgType = TEXT
		this.Content = "对不起，系统暂时不支持的类型"
	}
}

func ifBan(to_check string) bool {
	ban := strings.Split(beego.AppConfig.String("ban"), "&")
	for _, v := range ban {
		if strings.Contains(to_check, v) {
			return true
		}
	}
	return false
}

func ifLol(to_check string) bool {
	return strings.Contains(to_check, "lol#") || strings.Contains(to_check, "LOL#")
}

func ifMenu(to_check string) bool {
	return strings.Contains(to_check, "菜单") || strings.Contains(to_check, "？") || strings.Contains(to_check, "?")
}

func ifCalendar(to_check string) bool {
	return strings.Contains(to_check, "日历") || strings.Contains(to_check, "黄历")
}

func ifwxNews(to_check string) bool {
	return strings.Contains(to_check, "微信精选")
}

func ifNews(to_check string) bool {
	return strings.Contains(to_check, "新闻")
}

func getUser(uid string) *UserInfo {
	resp, err := http.Get("https://api.weixin.qq.com/cgi-bin/user/info?access_token=" + getAccessToken() + "&openid=" + uid + "&lang=zh_CN ")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body_bytes, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body_bytes),uid)
	user_info := new(UserInfo)
	if err := json.Unmarshal(body_bytes, user_info); err != nil {
		panic(err)
	}
	//fmt.Println(user_info)
	return user_info
}

func getAccessToken() string {
	resp, err := http.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" +
		beego.AppConfig.String("appID") + "&secret=" + beego.AppConfig.String("appsecret"))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body_bytes, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body_bytes))
	token := new(AccessToken)
	json.Unmarshal(body_bytes, token)
	return token.AccessToken
}