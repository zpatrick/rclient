package rclient

import (
	"net/http"
	"net/url"
)

type ClientOption func(*RestClient) error

func Client(client *http.Client) ClientOption {
	return func(r *RestClient) error {
		r.Client = client
		return nil
	}
}

func ReaderFAC(reader ReaderFactory) ClientOption {
	return func(r *RestClient) error {
		r.NewReader = reader
		return nil
	}
}

func SenderFAC(sender SenderFactory) ClientOption {
	return func(r *RestClient) error {
		r.NewSender = sender
		return nil
	}
}

func RequestOptions(options ...RequestOption) ClientOption {
	return func(r *RestClient) error {
		r.RequestOptions = append(r.RequestOptions, options...)
		return nil
	}
}

type RequestOption func(*http.Request) error

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
