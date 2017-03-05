package rclient

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestJSONReaderFactory(t *testing.T) {
	resp := &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString("\"body\"")),
	}

	var v string
	if err := JSONReaderFactory(&v)(resp); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "body", v)
}

func TestJSONReaderFactoryNilV(t *testing.T) {
	resp := &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString("\"body\"")),
	}

	if err := JSONReaderFactory(nil)(resp); err != nil {
		t.Fatal(err)
	}
}

func TestJSONReaderFactoryError_statusCode(t *testing.T) {
	codes := []int{0, 199, 300, 399, 400, 499, 500, 599}

	for _, c := range codes {
		resp := &http.Response{
			StatusCode: c,
			Body:       ioutil.NopCloser(bytes.NewBufferString("")),
		}

		if err := JSONReaderFactory(nil)(resp); err == nil {
			t.Fatalf("%d: Error was nil!", c)
		}
	}
}

func TestJSONReaderFactoryError_invalidJSON(t *testing.T) {
	resp := &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString("some_invalid_json")),
	}

	var p person
	if err := JSONReaderFactory(&p)(resp); err == nil {
		t.Fatalf("Error was nil!")
	}
}
