// Copyright (c) 2019 Cyro Dubeux. License MIT.

// Package tmdb is a wrapper for working with TMDb API.
package tmdb

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/avast/retry-go/v4"
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

var json = jsoniter.ConfigFastest

// TMDb constants
const (
	defaultBaseURL    = "https://api.themoviedb.org/3"
	alternateBaseURL  = "https://api.tmdb.org/3"
	permissionURL     = "https://www.themoviedb.org/authenticate/"
	authenticationURL = "/authentication/"
	movieURL          = "/movie/"
	tvURL             = "/tv/"
	tvSeasonURL       = "/season/"
	tvEpisodeURL      = "/episode/"
	personURL         = "/person/"
	searchURL         = "/search/"
	collectionURL     = "/collection/"
	companyURL        = "/company/"
	configurationURL  = "/configuration/"
	creditURL         = "/credit/"
	discoverURL       = "/discover/"
	networkURL        = "/network/"
	keywordURL        = "/keyword/"
	genreURL          = "/genre/"
	guestSessionURL   = "/guest_session/"
	listURL           = "/list/"
	accountURL        = "/account/"
	watchProvidersURL = "/watch/providers/"
)

var baseURL = alternateBaseURL

// Client type is a struct to instantiate this pkg.
type Client struct {
	// TMDb apiKey to use the client.
	apiKey string
	// bearerToken will be used for v4 requests.
	bearerToken string
	// sessionId to use the client.
	sessionID string
	// Auto retry flag to indicates if the client
	// should retry the previous operation.
	autoRetry bool
	// http.Client for custom configuration.
	// http    http.Client
	Timeout time.Duration
	req     *fasthttp.Request
	resp    *fasthttp.Response
}

// Response type is a struct for http responses.
type Response struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
}

// Init setups the Client with an apiKey.
func Init(apiKey string) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("api key is empty")
	}
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	return &Client{
		apiKey:    apiKey,
		req:       req,
		resp:      resp,
		sessionID: DemoSessionID,
	}, nil
}

// InitV4 setups the Client with an bearer token.
func InitV4(bearerToken string) (*Client, error) {
	if bearerToken == "" {
		return nil, errors.New("bearer token is empty")
	}
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	return &Client{
		bearerToken: bearerToken,
		req:         req,
		resp:        resp,
	}, nil
}

// SetSessionID will set the session id.
func (s *Client) SetSessionID(sid string) error {
	if sid == "" {
		return errors.New("the session id is empty")
	}
	s.sessionID = sid
	return nil
}

// SetClientConfig sets a custom configuration for the http.Client.
//func (c *Client) SetClientConfig(httpClient http.Client) {
//	c.http = httpClient
//}

// SetClientAutoRetry sets autoRetry flag to true.
func (s *Client) SetClientAutoRetry() {
	s.autoRetry = true
}

func (s *Client) Close() {
	if s.resp != nil {
		fasthttp.ReleaseResponse(s.resp)
	}
	if s.req != nil {
		fasthttp.ReleaseRequest(s.req)
	}
}

// Auto retry default duration.
const defaultRetryDuration = time.Second * 5

// retryDuration calculates the retry duration time.
func retryDuration(resp *http.Response) time.Duration {
	retryTime := resp.Header.Get("Retry-After")
	if retryTime == "" {
		return defaultRetryDuration
	}
	seconds, err := strconv.ParseInt(retryTime, 10, 32)
	if err != nil {
		return defaultRetryDuration
	}
	return time.Duration(seconds) * time.Second
}

// shouldRetry determines whether the status code indicates that the
// previous operation should be retried at a later time.
func shouldRetry(status int) bool {
	return status == http.StatusAccepted || status == http.StatusTooManyRequests
}

func (s *Client) get(url string, data interface{}) error {
	if url == "" {
		return errors.New("url field is empty")
	}
	if s.Timeout == 0 {
		s.Timeout = time.Second * 10
	}
	s.req.SetRequestURI(url)
	// s.req.Header.SetContentType("application/json")
	s.req.Header.Add("Accept", "application/json")
	// s.req.Header.SetMethod("POST")
	s.req.Header.SetMethod("GET")
	s.req.Header.Add("content-type", "application/json;charset=utf-8")
	if s.bearerToken != "" {
		s.req.Header.Add("Authorization", "Bearer "+s.bearerToken)
	}

	err := fasthttp.DoTimeout(s.req, s.resp, s.Timeout)
	if err != nil {
		return err
	}
	resBody := s.resp.Body()
	code := s.resp.StatusCode()

	if code == http.StatusTooManyRequests && s.autoRetry {
		// TODO:tsingson
		// return errors.New("too many requests")
		// ------------------------
		er1 := retry.Do(func() error {
			return fasthttp.DoTimeout(s.req, s.resp, s.Timeout)
		}, retry.Attempts(10), retry.Delay(time.Duration(15)*time.Second))
		if er1 != nil {
			return errors.New("too many requests")
		}
	}
	resBody = s.resp.Body()
	code = s.resp.StatusCode()
	//
	if code == http.StatusNoContent {
		return nil
	}
	if code != http.StatusOK {
		return s.decodeError(resBody, code)
	}
	buf := bytes.NewBuffer(resBody)
	if err = json.NewDecoder(buf).Decode(data); err != nil {
		return fmt.Errorf("could not decode the data: %s", err)
	}
	return nil
}

func (s *Client) request(
	url string,
	requestPayload interface{},
	method string,
	data interface{},
) error {
	if url == "" {
		return errors.New("url field is empty")
	}
	if s.Timeout == 0 {
		s.Timeout = time.Second * 10
	}
	bodyBytes := new(bytes.Buffer)
	json.NewEncoder(bodyBytes).Encode(requestPayload)

	s.req.SetRequestURI(url)
	// s.req.Header.SetContentType("application/json")
	s.req.Header.Add("Accept", "application/json")
	// s.req.Header.SetMethod("POST")
	s.req.Header.SetMethod(method)
	s.req.Header.Add("content-type", "application/json;charset=utf-8")
	if s.bearerToken != "" {
		s.req.Header.Add("Authorization", "Bearer "+s.bearerToken)
	}

	s.req.SetBody(bodyBytes.Bytes())

	err := fasthttp.DoTimeout(s.req, s.resp, s.Timeout)
	if err != nil {
		return fmt.Errorf("could not fetch the url: %s", err)
	}
	resBody := s.resp.Body()
	code := s.resp.StatusCode()
	if code == http.StatusTooManyRequests && s.autoRetry {
		// TODO:tsingson
		// return errors.New("too many requests")
		// ------------------------
		er1 := retry.Do(func() error {
			return fasthttp.DoTimeout(s.req, s.resp, s.Timeout)
		}, retry.Attempts(10), retry.Delay(time.Duration(15)*time.Second))
		if er1 != nil {
			return errors.New("too many requests")
		}
	}
	resBody = s.resp.Body()
	code = s.resp.StatusCode()
	//
	if code == http.StatusNoContent {
		return nil
	}
	if code != http.StatusOK {
		return s.decodeError(resBody, code)
	}
	buf := bytes.NewBuffer(resBody)
	if err = json.NewDecoder(buf).Decode(data); err != nil {
		return fmt.Errorf("could not decode the data: %s", err)
	}
	return nil
}

func (s *Client) fmtOptions(
	urlOptions map[string]string,
) string {
	options := ""
	if len(urlOptions) > 0 {
		for key, value := range urlOptions {
			options += fmt.Sprintf(
				"&%s=%s",
				key,
				url.QueryEscape(value),
			)
		}
	}
	return options
}

// SetAlternateBaseURL sets an alternate base url.
func (s *Client) SetAlternateBaseURL() {
	baseURL = alternateBaseURL
}

// SetCustomBaseURL sets an custom base url.
func (s *Client) SetCustomBaseURL(url string) {
	baseURL = url
}

// GetBaseURL gets the current base url.
func (s *Client) GetBaseURL() string {
	return baseURL
}

// Error type represents an error returned by the TMDB API.
type Error struct {
	StatusMessage string `json:"status_message,omitempty"`
	Success       bool   `json:"success,omitempty"`
	StatusCode    int    `json:"status_code,omitempty"`
}

func (e Error) Error() string {
	return fmt.Sprintf(
		"code: %d | success: %t | message: %s",
		e.StatusCode,
		e.Success,
		e.StatusMessage,
	)
}

func (s *Client) decodeError(resBody []byte, code int) error {
	if len(resBody) == 0 {
		return fmt.Errorf(
			"[%d]: empty body %s",
			code,
			http.StatusText(code),
		)
	}
	buf := bytes.NewBuffer(resBody)
	var clientError Error
	if err := json.NewDecoder(buf).Decode(&clientError); err != nil {
		return fmt.Errorf(
			"couldn't decode error: (%d) [%s]",
			len(resBody),
			resBody,
		)
	}
	return clientError
}
