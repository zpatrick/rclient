package rclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type RequestBuilder func(method, url string, body interface{}, options ...RequestOption) (*http.Request, error)

func BuildJSONRequest(method, url string, body interface{}, options ...RequestOption) (*http.Request, error) {
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

	return req, nil
}
