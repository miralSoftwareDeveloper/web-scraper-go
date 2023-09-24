package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly"
)

// initializing a data structure to keep the scraped data
type PokemonProduct struct {
	url, image, name, price string
}

func main() {
	// initializing the slice of structs to store the data to scrape
	var pokemonProducts []PokemonProduct

	// creating a new Colly instance
	c := colly.NewCollector()

	c.SetRequestTimeout(120 * time.Second)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})

	// visiting the target page

	// scraping logic
	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		pokemonProduct := PokemonProduct{}
		pokemonProduct.url = e.ChildAttr("a", "href")
		pokemonProduct.image = e.ChildAttr("img", "src")
		pokemonProduct.name = e.ChildText("h2")
		pokemonProduct.price = e.ChildText(".price")
		pokemonProducts = append(pokemonProducts, pokemonProduct)
		println("Scraping..", pokemonProduct.name)
	})
	// OnHTML function will not work if it is above Visit function.
	c.Visit("https://scrapeme.live/shop/")
	// opening the CSV file
	file, err := os.Create("products.csv")
	if err != nil {
		log.Fatalln("Failed to create output CSV file", err)
	}
	defer file.Close()

	// initializing a file writer
	writer := csv.NewWriter(file)

	// writing the CSV headers
	headers := []string{
		"url",
		"image",
		"name",
		"price",
	}
	writer.Write(headers)

	// writing each Pokemon product as a CSV row
	for _, pokemonProduct := range pokemonProducts {
		// converting a PokemonProduct to an array of strings
		record := []string{
			pokemonProduct.url,
			pokemonProduct.image,
			pokemonProduct.name,
			pokemonProduct.price,
		}

		// adding a CSV record to the output file
		writer.Write(record)
	}
	defer writer.Flush()

	fmt.Println("Press any key to exit...")
	_, err = fmt.Scanln()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
	}
	fmt.Println("Exit and Close...")
}
