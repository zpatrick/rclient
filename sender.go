package rclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Sender func() (*http.Response, error)

type SenderFactory func(client *http.Client, method, url string, body interface{}, options ...RequestOption) Sender

func JSONSenderFactory(client *http.Client, method, url string, body interface{}, options ...RequestOption) Sender {
	return func() (*http.Response, error) {
		b := new(bytes.Buffer)
		if err := json.NewEncoder(b).Encode(body); err != nil {
			return nil, err
		}

		req, err := http.NewRequest(method, url, b)
		if err != nil {
			return nil, err
		}

		req.Header.Add("content-type", "application/json")

		for _, option := range options {
			if err := option(req); err != nil {
				return nil, err
			}
		}

		return client.Do(req)
	}
}
