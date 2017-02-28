package rclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type RestClient struct {
	Host           string
	RequestDoer    RequestDoer
	RequestOptions []RequestOption
	Reader         Reader
}

func NewRestClient(host string, options ...ClientOption) (*RestClient, error) {
	r := &RestClient{
		Host:        host,
		RequestDoer: http.DefaultClient,
		Reader:      DefaultReader,
		RequestOptions: []RequestOption{
			func(req *http.Request) error {
				req.Header.Add("content-type", "application/json")
				return nil
			},
		},
	}

	for _, option := range options {
		if err := option(r); err != nil {
			return nil, err
		}
	}

	return r, nil
}

func (r *RestClient) Delete(path string, body, v interface{}, options ...RequestOption) error {
	return r.Do("DELETE", path, body, v, options...)
}

func (r *RestClient) Get(path string, v interface{}, options ...RequestOption) error {
	return r.Do("GET", path, nil, v, options...)
}

func (r *RestClient) Post(path string, body, v interface{}, options ...RequestOption) error {
	return r.Do("POST", path, body, v, options...)
}

func (r *RestClient) Put(path string, body, v interface{}, options ...RequestOption) error {
	return r.Do("PUT", path, body, v, options...)
}

func (r *RestClient) Do(method, path string, body, v interface{}, options ...RequestOption) error {
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s%s", r.Host, path)
	req, err := http.NewRequest(method, url, bytes.NewReader(b))
	if err != nil {
		return nil
	}

	for _, option := range append(r.RequestOptions, options...) {
		if err := option(req); err != nil {
			return err
		}
	}

	resp, err := r.RequestDoer.Do(req)
	if err != nil {
		return err
	}

	return r.Reader(resp, v)
}
