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
	// Enabling auto retry functionality.
	tmdbClient.SetClientAutoRetry()

	movie, err := tmdbClient.GetMovieDetails(299536, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(movie.Title)

	fmt.Println("------")

	// With options
	options := make(map[string]string)
	options["append_to_response"] = "credits"
	options["language"] = "pt-BR"

	movie, err = tmdbClient.GetMovieDetails(299536, options)
	if err != nil {
		fmt.Println(err)
	}

	// Credits - Iterate cast from append to response.
	for _, v := range movie.MovieCreditsAppend.Credits.Cast {
		fmt.Println(v.Name)
	}
}
