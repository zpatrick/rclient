package rclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type RequestSender interface {
	Send() (*http.Response, error)
}

type RequestSenderFunc func() (*http.Response, error)

func (s RequestSenderFunc) Send() (*http.Response, error) {
	return s()
}

type RequestSenderFactory func(client *http.Client, method, url string, body interface{}, options ...RequestOption) RequestSender

func JSONRequestSenderFactory(client *http.Client, method, url string, body interface{}, options ...RequestOption) RequestSender {
	return RequestSenderFunc(func() (*http.Response, error) {
		b := new(bytes.Buffer)
		if err := json.NewEncoder(b).Encode(body); err != nil {
			return nil, err
		}

		req, err := http.NewRequest(method, url, b)
		if err != nil {
			return nil, err
		}

		for _, option := range options {
			if err := option(req); err != nil {
				return nil, err
			}
		}

		return client.Do(req)
	})
}
