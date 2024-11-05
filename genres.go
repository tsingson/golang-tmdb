package tmdb

import "fmt"

// GenreMovieList type is a struct for genres movie list JSON response.
type GenreMovieList struct {
	Genres []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
}

// GetGenreMovieList get the list of official genres for movies.
//
// https://developers.themoviedb.org/3/genres/get-movie-list
func (s *Client) GetGenreMovieList(
	urlOptions map[string]string,
) (*GenreMovieList, error) {
	options := s.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%smovie/list?api_key=%s%s",
		baseURL,
		genreURL,
		s.apiKey,
		options,
	)
	genreMovieList := GenreMovieList{}
	if err := s.get(tmdbURL, &genreMovieList); err != nil {
		return nil, err
	}
	return &genreMovieList, nil
}

// GetGenreTVList get the list of official genres for TV shows.
//
// https://developers.themoviedb.org/3/genres/get-tv-list
func (s *Client) GetGenreTVList(
	urlOptions map[string]string,
) (*GenreMovieList, error) {
	options := s.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%stv/list?api_key=%s%s",
		baseURL,
		genreURL,
		s.apiKey,
		options,
	)
	genreTVList := GenreMovieList{}
	if err := s.get(tmdbURL, &genreTVList); err != nil {
		return nil, err
	}
	return &genreTVList, nil
}
