package core

import (
	"github.com/gocolly/colly/v2"
)

func ParseYears() []string {
	urls := []string{}
	// Instantiate default collector
	c := colly.NewCollector()

	// Extract product details
	c.OnHTML("a.grid_cell3", func(e *colly.HTMLElement) {
		urls = append(urls, e.Request.AbsoluteURL(e.Attr("href")))
	})

	c.Visit("https://www.kinoafisha.info/rating/movies/year/")

	return urls
}
