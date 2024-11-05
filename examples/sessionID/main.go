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

	// A valid session or guest session ID is required.
	//
	// You can read more about how this works:
	// https://developers.themoviedb.org/3/authentication/how-do-i-generate-a-session-id
	//
	// Once you have the SessionID, you can load it from a ENV variable or a database.
	if err = tmdbClient.SetSessionID(tmdb.DemoSessionID); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// PostMovieRating
	r, err := tmdbClient.PostMovieRating(299536, 3.5, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(r.StatusMessage)

	// DeleteMovieRating
	r, err = tmdbClient.DeleteMovieRating(299536, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(r.StatusMessage)
}
