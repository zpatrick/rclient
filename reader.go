package rclient

import (
	"encoding/json"
	"fmt"
	"github.com/zpatrick/go-series"
	"net/http"
)

type Reader func(resp *http.Response) error

type ReaderFactory func(v interface{}) Reader

func JSONReaderFactory(v interface{}) Reader {
	return func(resp *http.Response) error {
		defer resp.Body.Close()

		switch {
		case !series.Ints(200, 299).Contains(resp.StatusCode):
			return fmt.Errorf("Invalid status code: %d", resp.StatusCode)
		case v == nil:
			return nil
		default:
			return json.NewDecoder(resp.Body).Decode(v)
		}
	}
}
