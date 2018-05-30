package main

/*
<urlset>
	<url>
		<loc>http://www.thehindu.com/news/national/telangana/quality-education-message-on-wheels/article24027694.ece</loc>
		<news:news>
			<news:publication>
				<news:name>The Hindu</news:name>
				<news:language>en</news:language>
			</news:publication>
			<news:publication_date>2018-05-29T23:51:30+05:30</news:publication_date>
			<news:title>Quality education message on wheels</news:title>
			<news:keywords>Telangana State United Teachers Federation; TSUTF; jeep jatha; enrolment; dropout</news:keywords>
		</news:news>
	</url>
</urlset>
*/

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UrlsetIndex struct {
	Urls []Url `xml:"url"`
}

type Url struct {
	Location string `xml:"loc"`
}

func (u Url) String() string { // Add this function to any named type and you shall have a To string function
	return fmt.Sprintf(u.Location)
}

func main() {

	resp, _ := http.Get("http://www.thehindu.com/?service=newssitemap")
	bytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(bytes))

	var dataStruct UrlsetIndex
	xml.Unmarshal(bytes, &dataStruct)

	//fmt.Println(dataStruct.Urls)

	for _, url := range dataStruct.Urls { //index and value is given
		fmt.Printf("%s \n", url)
	}
}
