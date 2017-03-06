package rclient

import (
	"fmt"
	"net/http"
)

type RestClient struct {
	Host           string
	RequestBuilder RequestBuilder
	RequestDoer    RequestDoer
	ResponseReader ResponseReader
	RequestOptions []RequestOption
}

func NewRestClient(host string, options ...ClientOption) (*RestClient, error) {
	r := &RestClient{
		Host:           host,
		RequestBuilder: BuildJSONRequest,
		RequestDoer:    http.DefaultClient,
		ResponseReader: ReadJSONResponse,
		RequestOptions: []RequestOption{},
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
	url := fmt.Sprintf("%s%s", r.Host, path)
	options = append(r.RequestOptions, options...)

	req, err := r.RequestBuilder(method, url, body, options...)
	if err != nil {
		return err
	}

	resp, err := r.RequestDoer.Do(req)
	if err != nil {
		return err
	}

	return r.ResponseReader(resp, v)
}
