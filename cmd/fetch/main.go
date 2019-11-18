package main

import (
	"flag"
	"log"
	"os"

	"github.com/opencars/emblems/pkg/carlogos"
)

func main() {
	var path string

	flag.StringVar(&path, "path", "./emblems", "Path to save emblems")

	flag.Parse()

	client := carlogos.NewClient()

	// Create directory, if it is not exist.
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0755); err != nil {
			log.Fatal(err)
		}
	}

	// Scrape all emblems from the website.
	if err := client.ScrapeEmblems(path); err != nil {
		log.Fatal(err)
	}
}
