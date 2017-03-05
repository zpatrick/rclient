package rclient

import (
	"net/http"
)

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

type DoerFunc func(*http.Request) (*http.Response, error)

func (d DoerFunc) Do(req *http.Request) (*http.Response, error) {
	return d(req)
}
