package rclient

import (
	"net/http"
)

type RequestDoer interface {
	Do(*http.Request) (*http.Response, error)
}

type RequestDoerFunc func(*http.Request) (*http.Response, error)

func (d RequestDoerFunc) Do(req *http.Request) (*http.Response, error) {
	return d(req)
}
