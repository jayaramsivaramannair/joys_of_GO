package main

import (
	"fmt"
	"encoding/xml"
	"net/http"
	"io/ioutil"
)

var washPostXML = []byte (`
	<sitemapindex>
		<sitemap>
			<loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
		</sitemap>
		<sitemap>
			<loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
		</sitemap>
		<sitemap>
			<loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
		</sitemap>
	</sitemapindex>
	`)

type SitemapIndex struct {
	//The data type is a slice of location
	Locations [] string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword string
	Location string
}

func main() {
	//resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes := washPostXML
	var s SitemapIndex
	var n News

	news_map := make(map[string]NewsMap)

	xml.Unmarshal(bytes, &s)
	//fmt.Println(s.Locations)
	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
		for idx, _ := range n.Titles {
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}

	for idx, data := range news_map {
		fmt.Println("\n\n\n", idx)
		fmt.Println("\n", data.Keyword)
		fmt.Println("\n", data.Location)

	}
}