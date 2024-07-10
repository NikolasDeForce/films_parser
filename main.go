package main

import (
	"film/core"
)

func main() {
	urlsGenres := core.ParseGenres()
	urlsCountry := core.ParseCountry()
	urlsYears := core.ParseYears()

	core.ParseData(urlsGenres, "genres")
	core.ParseData(urlsCountry, "country")
	core.ParseData(urlsYears, "years")
}
