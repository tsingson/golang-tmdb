package main

import (
	"fmt"
	"os"

	tmdb "github.com/cyruzin/golang-tmdb"
)

func main() {
	tmdbClient, err := tmdb.Init(tmdb.DemoApiKey)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	trending, err := tmdbClient.GetTrending("movie", "week", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, result := range trending.Results {
		fmt.Println(result.Title)
	}

	fmt.Println("------")

	// With options
	options := make(map[string]string)
	options["page"] = "2"
	options["language"] = "es-ES"

	trending, err = tmdbClient.GetTrending("tv", "day", options)
	if err != nil {
		fmt.Println(err)
	}

	for _, result := range trending.Results {
		fmt.Println(result.Name)
	}
}
