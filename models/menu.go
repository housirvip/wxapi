package models

import (
	"strings"
	"io/ioutil"
	"net/http"
)

func CreateMenu() string {
	menu:=` {
     "button":[
     {
          "type":"click",
          "name":"黄历",
          "key":"黄历"
      },
      {
           "name":"菜单",
           "sub_button":[
           {
               "type":"view",
               "name":"搜索",
               "url":"http://www.soso.com/"
            },
            {
               "type":"view",
               "name":"视频",
               "url":"http://v.qq.com/"
            },
            {
               "type":"click",
               "name":"赞一下我们",
               "key":"V1001_GOOD"
            }]
       }]
 }`
	access_tocken:=getAccessToken()
	resp, _ := http.Post("https://api.weixin.qq.com/cgi-bin/menu/create?access_token="+access_tocken, "application/json",
		strings.NewReader(menu))
	defer resp.Body.Close()
	body_bytes, _ := ioutil.ReadAll(resp.Body)
	return string(body_bytes)
}

func DeleteMenu() string {
	access_tocken:=getAccessToken()
	resp, _ := http.Get("https://api.weixin.qq.com/cgi-bin/menu/delete?access_token="+access_tocken)
	defer resp.Body.Close()
	body_bytes, _ := ioutil.ReadAll(resp.Body)
	return string(body_bytes)

}

func GetMenu() string {
	access_tocken:=getAccessToken()
	resp, _ := http.Get("https://api.weixin.qq.com/cgi-bin/menu/get?access_token="+access_tocken)
	defer resp.Body.Close()
	body_bytes, _ := ioutil.ReadAll(resp.Body)
	return string(body_bytes)

}
