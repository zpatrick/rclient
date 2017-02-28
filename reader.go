package rclient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Reader func(*http.Response, interface{}) error

func DefaultReader(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("Invalid status code: %d", resp.StatusCode)
	}

	if v == nil {
		return nil
	}

	return json.NewDecoder(resp.Body).Decode(v)
}
