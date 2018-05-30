package main

import (
	"encoding/xml"
	"html/template"
	"io/ioutil"
	"net/http"
)

type News struct {
	Urls     []string `xml:"url>loc"`
	Titles   []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
}

type NewsMap struct {
	Url     string
	Keyword string
}

type NewsMapAgg struct {
	Title string
	News  map[string]NewsMap
}

func newsAggregatorHandler(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get("http://www.thehindu.com/?service=newssitemap")
	bytes, _ := ioutil.ReadAll(resp.Body)
	t, _ := template.ParseFiles("template.html")

	var data News
	xml.Unmarshal(bytes, &data)
	newsmap := make(map[string]NewsMap)

	for i, v := range data.Titles {
		newsmap[v] = NewsMap{Url: data.Urls[i], Keyword: data.Keywords[i]}
	}

	p := NewsMapAgg{Title: "News Aggregator", News: newsmap}
	t.Execute(w, p)

}

func InjectAndServe() {
	defer http.ListenAndServe(":3000", nil)
	http.HandleFunc("/", newsAggregatorHandler)
}

func main() {
	InjectAndServe()
}
