package tmdb

import (
	"fmt"
	"net/url"
)

// SearchCompanies type is a struct for companies JSON response.
type SearchCompanies struct {
	Page int64 `json:"page"`
	*SearchCompaniesResults
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
}

// GetSearchCompanies search for companies.
//
// https://developers.themoviedb.org/3/search/search-companies
func (s *Client) GetSearchCompanies(
	query string,
	urlOptions map[string]string,
) (*SearchCompanies, error) {
	options := s.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%scompany?api_key=%s&query=%s%s",
		baseURL,
		searchURL,
		s.apiKey,
		url.QueryEscape(query),
		options,
	)
	SearchCompanies := SearchCompanies{}
	if err := s.get(tmdbURL, &SearchCompanies); err != nil {
		return nil, err
	}
	return &SearchCompanies, nil
}

// SearchCollections type is a strcut for collections JSON response.
type SearchCollections struct {
	Page int64 `json:"page"`
	*SearchCollectionsResults
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
}

// GetSearchCollections search for collections.
//
// https://developers.themoviedb.org/3/search/search-collections
func (s *Client) GetSearchCollections(
	query string,
	urlOptions map[string]string,
) (*SearchCollections, error) {
	options := s.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%scollection?api_key=%s&query=%s%s",
		baseURL,
		searchURL,
		s.apiKey,
		url.QueryEscape(query),
		options,
	)
	searchCollections := SearchCollections{}
	if err := s.get(tmdbURL, &searchCollections); err != nil {
		return nil, err
	}
	return &searchCollections, nil
}

// SearchKeywords type is a struct for keywords JSON response.
type SearchKeywords struct {
	Page int64 `json:"page"`
	*SearchKeywordsResults
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
}

// GetSearchKeywords search for keywords.
//
// https://developers.themoviedb.org/3/search/search-keywords
func (s *Client) GetSearchKeywords(
	query string,
	urlOptions map[string]string,
) (*SearchKeywords, error) {
	options := s.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%skeyword?api_key=%s&query=%s%s",
		baseURL,
		searchURL,
		s.apiKey,
		url.QueryEscape(query),
		options,
	)
	searchKeywords := SearchKeywords{}
	if err := s.get(tmdbURL, &searchKeywords); err != nil {
		return nil, err
	}
	return &searchKeywords, nil
}

// SearchMovies type is a struct for movies JSON response.
type SearchMovies struct {
	Page         int64 `json:"page"`
	TotalResults int64 `json:"total_results"`
	TotalPages   int64 `json:"total_pages"`
	*SearchMoviesResults
}

// GetSearchMovies search for keywords.
//
// https://developers.themoviedb.org/3/search/search-movies
func (s *Client) GetSearchMovies(
	query string,
	urlOptions map[string]string,
) (*SearchMovies, error) {
	options := s.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%smovie?api_key=%s&query=%s%s",
		baseURL,
		searchURL,
		s.apiKey,
		url.QueryEscape(query),
		options,
	)
	searchMovies := SearchMovies{}
	if err := s.get(tmdbURL, &searchMovies); err != nil {
		return nil, err
	}
	return &searchMovies, nil
}

// SearchMulti type is a struct for multi JSON response.
type SearchMulti struct {
	Page int `json:"page"`
	*SearchMultiResults
	TotalResults int64 `json:"total_results"`
	TotalPages   int64 `json:"total_pages"`
}

// GetSearchMulti search multiple models in a single request.
// Multi search currently supports searching for movies,
// tv shows and people in a single request.
//
// https://developers.themoviedb.org/3/search/multi-search
func (s *Client) GetSearchMulti(
	query string,
	urlOptions map[string]string,
) (*SearchMulti, error) {
	options := s.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%smulti?api_key=%s&query=%s%s",
		baseURL,
		searchURL,
		s.apiKey,
		url.QueryEscape(query),
		options,
	)
	searchMulti := SearchMulti{}
	if err := s.get(tmdbURL, &searchMulti); err != nil {
		return nil, err
	}
	return &searchMulti, nil
}

// SearchPeople type is a struct for people JSON response.
type SearchPeople struct {
	Page         int64 `json:"page"`
	TotalResults int64 `json:"total_results"`
	TotalPages   int64 `json:"total_pages"`
	*SearchPeopleResults
}

// GetSearchPeople search for people.
//
// https://developers.themoviedb.org/3/search/search-people
func (s *Client) GetSearchPeople(
	query string,
	urlOptions map[string]string,
) (*SearchPeople, error) {
	options := s.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%sperson?api_key=%s&query=%s%s",
		baseURL,
		searchURL,
		s.apiKey,
		url.QueryEscape(query),
		options,
	)
	searchPeople := SearchPeople{}
	if err := s.get(tmdbURL, &searchPeople); err != nil {
		return nil, err
	}
	return &searchPeople, nil
}

// SearchTVShows type is a struct for tv show JSON response.
type SearchTVShows struct {
	Page         int64 `json:"page"`
	TotalResults int64 `json:"total_results"`
	TotalPages   int64 `json:"total_pages"`
	*SearchTVShowsResults
}

// GetSearchTVShow search for a TV Show.
//
// https://developers.themoviedb.org/3/search/search-tv-shows
func (s *Client) GetSearchTVShow(
	query string,
	urlOptions map[string]string,
) (*SearchTVShows, error) {
	options := s.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%stv?api_key=%s&query=%s%s",
		baseURL,
		searchURL,
		s.apiKey,
		url.QueryEscape(query),
		options,
	)
	searchTVShows := SearchTVShows{}
	if err := s.get(tmdbURL, &searchTVShows); err != nil {
		return nil, err
	}
	return &searchTVShows, nil
}
