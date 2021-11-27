package main

import (
	"fmt"
	"encoding/xml"
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
	Locations []Location `xml:"sitemap"`
}


type Location struct {
	Loc string `xml:"loc"`
}


func(l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	//resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes := washPostXML
	var s SitemapIndex
	xml.Unmarshal(bytes, &s)
	fmt.Println(s.Locations)
}