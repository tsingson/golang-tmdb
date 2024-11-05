package tmdb

import "fmt"

// ChangesMovie type is a struct for movie changes JSON response.
type ChangesMovie struct {
	*ChangesMovieResults
	Page         int64 `json:"page"`
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
}

// GetChangesMovie get a list of all of the movie ids
// that have been changed in the past 24 hours.
//
// You can query it for up to 14 days worth of changed IDs
// at a time with the start_date and end_date query parameters.
// 100 items are returned per page.
//
// https://developers.themoviedb.org/3/changes/get-movie-change-list
func (s *Client) GetChangesMovie(
	urlOptions map[string]string,
) (*ChangesMovie, error) {
	options := s.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%schanges?api_key=%s%s",
		baseURL,
		movieURL,
		s.apiKey,
		options,
	)
	changesMovies := ChangesMovie{}
	if err := s.get(tmdbURL, &changesMovies); err != nil {
		return nil, err
	}
	return &changesMovies, nil
}

// ChangesTV type is a struct for tv changes JSON response.
type ChangesTV struct {
	*ChangesMovie
}

// GetChangesTV get a list of all of the TV show ids
// that have been changed in the past 24 hours.
//
// You can query it for up to 14 days worth of changed IDs
// at a time with the start_date and end_date query parameters.
// 100 items are returned per page.
//
// https://developers.themoviedb.org/3/changes/get-tv-change-list
func (s *Client) GetChangesTV(
	urlOptions map[string]string,
) (*ChangesTV, error) {
	options := s.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%schanges?api_key=%s%s",
		baseURL,
		tvURL,
		s.apiKey,
		options,
	)
	changesTV := ChangesTV{}
	if err := s.get(tmdbURL, &changesTV); err != nil {
		return nil, err
	}
	return &changesTV, nil
}

// ChangesPerson type is a struct for person changes JSON response.
type ChangesPerson struct {
	*ChangesMovie
}

// GetChangesPerson get a list of all of the person ids
// that have been changed in the past 24 hours.
//
// You can query it for up to 14 days worth of changed IDs
// at a time with the start_date and end_date query parameters.
// 100 items are returned per page.
//
// https://developers.themoviedb.org/3/changes/get-person-change-list
func (s *Client) GetChangesPerson(
	urlOptions map[string]string,
) (*ChangesPerson, error) {
	options := s.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%schanges?api_key=%s%s",
		baseURL,
		personURL,
		s.apiKey,
		options,
	)
	changesPerson := ChangesPerson{}
	if err := s.get(tmdbURL, &changesPerson); err != nil {
		return nil, err
	}
	return &changesPerson, nil
}
