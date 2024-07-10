package core

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

func ParseData(urls []string, from string) {
	for i := 0; i < len(urls); i++ {
		lastIndex := strings.Split(urls[i], "/")
		fName := fmt.Sprintf("./core/dates/%s/%v.csv", from, lastIndex[len(lastIndex)-2])
		file, err := os.Create(fName)
		if err != nil {
			log.Fatalf("Cannot create file %v: %s\n", file, err)
		}
		defer file.Close()
		writer := csv.NewWriter(file)
		defer writer.Flush()
		// Write CSV header
		writer.Write([]string{"Rating", "Name", "Genres", "Year/Country", "URL"})

		// Instantiate default collector
		c := colly.NewCollector()

		// Extract product details
		c.OnHTML("div.movieItem", func(e *colly.HTMLElement) {
			writer.Write([]string{
				e.ChildText("span.movieItem_itemRating"),
				e.ChildText("a.movieItem_title"),
				e.ChildText("span.movieItem_genres"),
				e.ChildText("span.movieItem_year"),
				e.ChildAttr("a", "href"),
			})
		})

		//Find and visit next page links
		c.OnHTML("nav.ratings_pagination", func(e *colly.HTMLElement) {
			lastQuery := e.ChildAttrs("a", "href")
			e.Request.Visit(lastQuery[len(lastQuery)-1])
		})

		c.Visit(urls[i])

		log.Printf("Scraping %s", urls[i])
	}

	log.Println("Scraping finished, check the files!")

}
