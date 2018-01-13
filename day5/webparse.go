package main

/*
<sitemapindex>
	<sitemap>
		<loc></loc>
	</sitemap>
</sitemapindex>
*/

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SitemapnIndex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func main() {
	resp, _ := http.Get("the url goes here")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	var s SitemapnIndex
	xml.Unmarshal(bytes, &s)
	fmt.Println(s.Locations)
}
