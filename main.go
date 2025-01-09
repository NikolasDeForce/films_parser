package main

import (
	"film/core"
	"fmt"
	"log"
	"sync"
)

func main() {
	//get urls from categories
	urlsGenres := core.ParseGenres()
	urlsCountry := core.ParseCountry()
	urlsYears := core.ParseYears()

	// Собраем данные паралельно горутинами со всех категорий и сохраняем в папки
	var wg sync.WaitGroup

	for _, date := range urlsGenres {
		fmt.Printf("goroutine start %v\n", date)

		wg.Add(1)

		go func(page string) {
			defer wg.Done()
			core.ParseData(page, "genres")
			fmt.Printf("goroutine finish %v\n", page)
			fmt.Println("----------------------------------")
		}(date)
	}

	for _, date := range urlsCountry {
		fmt.Printf("goroutine start %v\n", date)

		wg.Add(1)

		go func(page string) {
			defer wg.Done()
			fmt.Println("----------------------------------")
			core.ParseData(page, "country")
			fmt.Printf("goroutine finish %v\n", page)
			fmt.Println("----------------------------------")
		}(date)
	}

	for _, date := range urlsYears {
		fmt.Printf("goroutine start %v\n", date)

		wg.Add(1)

		go func(page string) {
			defer wg.Done()
			core.ParseData(page, "years")
			fmt.Printf("goroutine finish %v\n", page)
			fmt.Println("----------------------------------")
		}(date)
	}

	wg.Wait()

	log.Println("Scraping finished, check the files!")
	// scraping data
	// core.ParseData(urlsGenres, "genres")
	// core.ParseData(urlsCountry, "country")
	// core.ParseData(urlsYears, "years")
}
