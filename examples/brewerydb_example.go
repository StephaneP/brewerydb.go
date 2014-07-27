package main

import (
	"github.com/stephanep/brewerydb.go"
	"log"
	"os"
)

func main() {
	client := brewerydb.NewClient(os.Getenv("BREWERYDB_KEY"))

	beer, err := client.GetBeer("tmEthz")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Beer", beer)
}
