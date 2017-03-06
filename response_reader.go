package rclient

import (
	"encoding/json"
	"fmt"
	"github.com/zpatrick/go-series"
	"net/http"
)

type ResponseReader func(resp *http.Response, v interface{}) error

func ReadJSONResponse(resp *http.Response, v interface{}) error {
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
