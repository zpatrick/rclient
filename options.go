package rclient

import (
	"net/http"
	"net/url"
)

type ClientOption func(client *RestClient) error

func Builder(builder RequestBuilder) ClientOption {
	return func(r *RestClient) error {
		r.RequestBuilder = builder
		return nil
	}
}

func Doer(doer RequestDoer) ClientOption {
	return func(r *RestClient) error {
		r.RequestDoer = doer
		return nil
	}
}

func Reader(reader ResponseReader) ClientOption {
	return func(r *RestClient) error {
		r.ResponseReader = reader
		return nil
	}
}

func RequestOptions(options ...RequestOption) ClientOption {
	return func(r *RestClient) error {
		r.RequestOptions = append(r.RequestOptions, options...)
		return nil
	}
}

type RequestOption func(req *http.Request) error

func BasicAuth(user, pass string) RequestOption {
	return func(req *http.Request) error {
		req.SetBasicAuth(user, pass)
		return nil
	}
}

func Header(name, val string) RequestOption {
	return func(req *http.Request) error {
		req.Header.Add(name, val)
		return nil
	}
}

func Headers(headers map[string]string) RequestOption {
	return func(req *http.Request) error {
		for name, val := range headers {
			req.Header.Add(name, val)
		}

		return nil
	}
}

func Query(query url.Values) RequestOption {
	return func(req *http.Request) error {
		req.URL.RawQuery = query.Encode()
		return nil
	}
}
