package rclient

import (
	"fmt"
	"net/http"
)

type RestClient struct {
	Host           string
	Client         *http.Client
	NewReader         ResponseReaderFactory
	NewSender         RequestSenderFactory
	RequestOptions []RequestOption
}

func NewRestClient(host string, options ...ClientOption) (*RestClient, error) {
	r := &RestClient{
		Host:   host,
		Client: http.DefaultClient,
		NewReader: JSONResponseReaderFactory,
		NewSender: JSONRequestSenderFactory,
		RequestOptions: []RequestOption{
			Header("content-type", "application/json"),
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
	url := fmt.Sprintf("%s%s", r.Host, path)
	options = append(r.RequestOptions, options...)
	sender := r.NewSender(r.Client, method, url, body, options...)
	reader := r.NewReader(v)
	return Do(sender, reader)
}
