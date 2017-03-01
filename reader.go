package rclient

import (
	"encoding/json"
	"fmt"
	"github.com/zpatrick/go-series"
	"net/http"
)

type ResponseReader interface {
	Read(*http.Response) error
}

type ResponseReaderFunc func(*http.Response) error

func (r ResponseReaderFunc) Read(resp *http.Response) error {
	return r(resp)
}

type ResponseReaderFactory func(v interface{}) ResponseReader

func JSONResponseReaderFactory(v interface{}) ResponseReader {
	return ResponseReaderFunc(func(resp *http.Response) error {
		defer resp.Body.Close()

		switch {
		case !series.Ints(200, 299).Contains(resp.StatusCode):
			return fmt.Errorf("Invalid status code: %d", resp.StatusCode)
		case v == nil:
			return nil
		default:
			return json.NewDecoder(resp.Body).Decode(v)
		}
	})
}
