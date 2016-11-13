package url2io

import (
	"net/http"
	"io/ioutil"
	"github.com/bitly/go-simplejson"
)

func API(url string, token string, text string, next string) (getinfo []string) {
	reqest, _ := http.Get("http://api.url2io.com/article?token=" + token + "&url=" + url + "&fields=" + text + next)
result, _ := ioutil.ReadAll(reqest.Body)
str,_:=simplejson.NewJson(result)
title, _ := str.Get("title").String()
urls,_:=str.Get("url").String()
date,_:=str.Get("date").String()
	if text != nil {
		texts, _ := str.Get("text").String()
		getinfo = append(getinfo, texts)
	}
	if text == nil {
		content := str.Get("content")
		getinfo = append(getinfo, content)
	}
	if next != nil {
		nexts, _ := str.Get("next").String()
		getinfo = append(getinfo, nexts)
	}
	if next == nil {}
	getinfo = append(getinfo, urls, date, title)
	return getinfo
}
