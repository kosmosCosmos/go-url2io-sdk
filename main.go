package chuansongmen

import (
	"net/http"
	"io/ioutil"
	"github.com/bitly/go-simplejson"
)

func API(url string,token string)(getinfo []string) {
reqest, _ := http.Get("http://api.url2io.com/article?token="+token+"&url="+url+"&fields=text")
result, _ := ioutil.ReadAll(reqest.Body)
str,_:=simplejson.NewJson(result)
title, _ := str.Get("title").String()
urls,_:=str.Get("url").String()
date,_:=str.Get("date").String()
text,_:=str.Get("text").String()
next,_:=str.Get("next").String()
getinfo=append(getinfo,urls,date,title,text,next)
return getinfo
}
