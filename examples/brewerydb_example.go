package main

import (
	"github.com/stephanep/brewerydb.go"
	"log"
)

func main() {
	client := brewerydb.NewClient("<api_key>")

	beer, err := client.GetBeer("<beer_id>")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Beer", beer)
}
