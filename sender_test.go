package rclient

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestJSONSenderFactory(t *testing.T) {
	doer := DoerFunc(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "www.domain.com/path", req.URL.String())
		assert.Equal(t, "application/json", req.Header.Get("content-type"))

		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "\"body\"\n", string(body))

		return nil, nil
	})

	send := JSONSenderFactory(doer, "GET", "www.domain.com/path", "body")
	if _, err := send(); err != nil {
		t.Fatal(err)
	}
}

func TestJSONSenderFactoryOptions(t *testing.T) {
	t.Skip("todo")
}

func TestJSONSenderFactoryError(t *testing.T) {
	doer := DoerFunc(func(req *http.Request) (*http.Response, error) {
		return nil, errors.New("some_error")
	})

	send := JSONSenderFactory(doer, "GET", "www.domain.com/path", "body")
	if _, err := send(); err == nil {
		t.Fatalf("Error was nil!")
	}
}
