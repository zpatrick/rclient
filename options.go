package rclient

import (
	"net/http"
)

type ClientOption func(*RestClient) error

func WithRequestDoer(doer RequestDoer) ClientOption {
	return func(r *RestClient) error {
		r.RequestDoer = doer
		return nil
	}
}

func WithReader(reader func(*http.Response, interface{}) error) ClientOption {
	return func(r *RestClient) error {
		r.Reader = reader
		return nil
	}
}

func WithBasicAuth(user, pass string) ClientOption {
	return func(r *RestClient) error {
		r.RequestOptions = append(r.RequestOptions, ReqWithBasicAuth(user, pass))
		return nil
	}
}

func WithHeader(name, val string) func(r *RestClient) error {
	return func(r *RestClient) error {
		r.RequestOptions = append(r.RequestOptions, ReqWithHeader(name, val))
		return nil
	}
}

func WithHeaders(headers map[string]string) ClientOption {
	return func(r *RestClient) error {
		r.RequestOptions = append(r.RequestOptions, ReqWithHeaders(headers))
		return nil
	}
}

type RequestOption func(*http.Request) error

func ReqWithBasicAuth(user, pass string) RequestOption {
	return func(req *http.Request) error {
		req.SetBasicAuth(user, pass)
		return nil
	}
}

func ReqWithHeader(name, val string) RequestOption {
	return func(req *http.Request) error {
		req.Header.Add(name, val)
		return nil
	}
}

func ReqWithHeaders(headers map[string]string) RequestOption {
	return func(req *http.Request) error {
		for name, val := range headers {
			req.Header.Add(name, val)
		}

		return nil
	}
}

func ReqWithQuery(params map[string]string) RequestOption {
	return func(req *http.Request) error {
		return nil
	}
}
