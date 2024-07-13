package main

import (
	"film/core"
)

func main() {
	//get urls from categories
	urlsGenres := core.ParseGenres()
	urlsCountry := core.ParseCountry()
	urlsYears := core.ParseYears()

	// scraping data
	core.ParseData(urlsGenres, "genres")
	core.ParseData(urlsCountry, "country")
	core.ParseData(urlsYears, "years")
}
