package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

type News struct {
	Urls     []string `xml:"url>loc"`
	Titles   []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
}

type NewsMap struct {
	Url      string
	Keywords string
}

type NewsMapAgg struct {
	Title string
	News  map[string]NewsMap
}

func newsRoutine(loc string, c chan News) {
	defer wg.Done()
	resp, _ := http.Get(loc)
	bytes, _ := ioutil.ReadAll(resp.Body)

	var news News
	xml.Unmarshal(bytes, &news)
	resp.Body.Close()
	c <- news
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	queue := make(chan News, 1000)

	wg.Add(1)
	go newsRoutine("http://www.thehindu.com/?service=newssitemap", queue)
	wg.Wait()
	close(queue)

	newsmap := make(map[string]NewsMap)
	for n := range queue {
		for i, v := range n.Titles {
			newsmap[v] = NewsMap{Url: n.Urls[i], Keywords: n.Keywords[i]}
		}
	}

	newsmapagg := NewsMapAgg{Title: "News Aggregator", News: newsmap}
	t, _ := template.ParseFiles("template.html")
	t.Execute(w, newsmapagg)

}

func main() {

	fmt.Println("News Aggregator")

	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":3000", nil)

}
