package rclient

import (
	"net/http"
)

type RequestDoer interface {
	Do(*http.Request) (*http.Response, error)
}

type DoerFunc func(*http.Request) (*http.Response, error)

func (do DoerFunc) Do(req *http.Request) (*http.Response, error) {
	return do(req)
}
