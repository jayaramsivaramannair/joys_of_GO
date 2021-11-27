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
	Keywords []string `xml:"url>news>Keywords"`
	Locations []string `xml:"url>loc"`
}

func main() {
	//resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes := washPostXML
	var s SitemapIndex
	var n News
	xml.Unmarshal(bytes, &s)
	//fmt.Println(s.Locations)
	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
	}

	fmt.Println(n)
}