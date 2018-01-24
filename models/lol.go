package models

import (
	"github.com/PuerkitoBio/goquery"
	"strconv"
)

const (
	暂无 = iota
	艾欧尼亚
	比尔吉沃特
	祖安
	诺克萨斯
	班德尔城
	德玛西亚
	皮尔特沃夫
	战争学院
	弗雷尔卓德
	巨神峰
	雷瑟守备
	无畏先锋
	裁决之地
	黑色玫瑰
	暗影岛
	恕瑞玛
	钢铁烈阳
	水晶之痕
	均衡教派
	扭曲丛林
	教育网专区
	影流
	守望之海
	征服之海
	卡拉曼达
	巨龙之巢
	皮城警备
)

//type Server struct {
//	Id int
//	Name string
//	Key int
//} 

type LolInfo struct {
	Server string
	Gameid string
	Info   string
}

func (this *LolInfo) GetInfo() string {
	server_id := matchServer(this.Server)
	//fmt.Println(server_id,this)
	if server_id == 0 {
		return "您所填写的大区找不到"
	}
	to_search := "http://lol.766.com/gamer/" + strconv.Itoa(server_id) + "/" + this.Gameid
	//fmt.Println(to_search)
	doc, err := goquery.NewDocument(to_search)
	if err != nil {
		panic(err)
		return "服务器内部错误"
	}
	//fmt.Println(doc.Html())
	if len(doc.Find(".name").Nodes) == 0 {
		return "您所填写的游戏ID找不到"
	}
	this.Info = "游戏ID:" + this.Gameid +
		"\n大区:" + this.Server +
		"\n段位:" + doc.Find(".level").Find(".value").Text() +
		"\n胜点:" + doc.Find(".winpoint").Find(".value").Text() +
		"\n近30天\n胜场:" + doc.Find(".win").Text() +
		"\n负场:" + doc.Find(".lose").Text() +
		"\n" + doc.Find(".value.winrate").Text() +
		"\n平均KDA:" + doc.Find(".kda").Eq(0).Find(".value").Text()
	return this.Info
}

func matchServer(server string) int {
	switch server {
	case "艾欧尼亚":return 艾欧尼亚
	case "比尔吉沃特":return 比尔吉沃特
	case "祖安":return 祖安
	case "诺克萨斯":return 诺克萨斯
	case "班德尔城":return 班德尔城
	case "德玛西亚":return 德玛西亚
	case "皮尔特沃夫":return 皮尔特沃夫
	case "战争学院":return 战争学院
	case "弗雷尔卓德":return 弗雷尔卓德
	case "巨神峰":return 巨神峰
	case "雷瑟守备":return 雷瑟守备
	case "无畏先锋":return 无畏先锋
	case "裁决之地":return 裁决之地
	case "黑色玫瑰":return 黑色玫瑰
	case "暗影岛":return 暗影岛
	case "恕瑞玛":return 恕瑞玛
	case "钢铁烈阳":return 钢铁烈阳
	case "水晶之痕":return 水晶之痕
	case "均衡教派":return 均衡教派
	case "扭曲丛林":return 扭曲丛林
	case "教育网专区":return 教育网专区
	case "影流":return 影流
	case "守望之海":return 守望之海
	case "征服之海":return 征服之海
	case "卡拉曼达":return 卡拉曼达
	case "巨龙之巢":return 巨龙之巢
	case "皮城警备":return 皮城警备
	default:
		return 暂无
	}
}