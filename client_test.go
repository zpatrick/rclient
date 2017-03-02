package rclient

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestClientDelete(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method)
		assert.Equal(t, "/people/john", r.URL.Path)

		write(t, w, 200, nil)
	}

	client, server := newClientAndServer(t, handler)
	defer server.Close()

	if err := client.Delete("/people/john", nil, nil); err != nil {
		t.Error(err)
	}
}

func TestClientDeleteWithBody(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method)
		assert.Equal(t, "/people", r.URL.Path)

		var p person
		read(t, r, &p)

		assert.Equal(t, person{"John Doe", 30}, p)

		write(t, w, 200, p)
	}

	client, server := newClientAndServer(t, handler)
	defer server.Close()

	var p person
	request := person{Name: "John Doe", Age: 30}
	if err := client.Delete("/people", request, &p); err != nil {
		t.Error(err)
	}

	assert.Equal(t, person{"John Doe", 30}, p)
}

func TestClientGet(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/people/john", r.URL.Path)

		write(t, w, 200, person{Name: "John Doe", Age: 30})
	}

	client, server := newClientAndServer(t, handler)
	defer server.Close()

	var p person
	if err := client.Get("/people/john", &p); err != nil {
		t.Error(err)
	}

	assert.Equal(t, person{"John Doe", 30}, p)
}

func TestClientPost(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/people", r.URL.Path)

		var p person
		read(t, r, &p)

		assert.Equal(t, person{"John Doe", 30}, p)

		write(t, w, 201, p)
	}

	client, server := newClientAndServer(t, handler)
	defer server.Close()

	var p person
	request := person{Name: "John Doe", Age: 30}
	if err := client.Post("/people", request, &p); err != nil {
		t.Error(err)
	}

	assert.Equal(t, person{"John Doe", 30}, p)
}

func TestClientPut(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PUT", r.Method)
		assert.Equal(t, "/people/john", r.URL.Path)

		var p person
		read(t, r, &p)

		assert.Equal(t, person{"John Do", 35}, p)

		write(t, w, 200, p)
	}

	client, server := newClientAndServer(t, handler)
	defer server.Close()

	var p person
	request := person{Name: "John Do", Age: 35}
	if err := client.Put("/people/john", request, &p); err != nil {
		t.Error(err)
	}

	assert.Equal(t, person{"John Do", 35}, p)
}

func TestClientDo(t *testing.T) {
	sFactory := func(c *http.Client, method, url string, body interface{}, options ...RequestOption) Sender {
		return func() (*http.Response, error) {
			assert.Equal(t, "POST", method)
			assert.Equal(t, "https://domain.com/path", url)
			assert.Equal(t, "body", body)
			assert.Len(t, options, 0)

			return nil, nil
		}
	}

	p := person{Name: "John Doe"}
	rFactory := func(v interface{}) Reader {
		return func(*http.Response) error {
			assert.Equal(t, &p, v)

			return nil
		}
	}

	client, err := NewRestClient("https://domain.com", SenderFAC(sFactory), ReaderFAC(rFactory))
	if err != nil {
		t.Fatal(err)
	}

	if err := client.Post("/path", "body", &p); err != nil {
		t.Fatal(err)
	}
}

func TestClientSenderError(t *testing.T) {
	sFactory := func(*http.Client, string, string, interface{}, ...RequestOption) Sender {
		return func() (*http.Response, error) {
			return nil, errors.New("some error")
		}
	}

	client, err := NewRestClient("", SenderFAC(sFactory))
	if err != nil {
		t.Fatal(err)
	}

	if err := client.Get("/path", nil); err == nil {
		t.Fatal("Error was nil!")
	}
}

func TestClientReaderError(t *testing.T) {
	sFactory := func(*http.Client, string, string, interface{}, ...RequestOption) Sender {
		return func() (*http.Response, error) {
			return nil, nil
		}
	}

	rFactory := func(interface{}) Reader {
		return func(*http.Response) error {
			return errors.New("some error")
		}
	}

	client, err := NewRestClient("", SenderFAC(sFactory), ReaderFAC(rFactory))
	if err != nil {
		t.Fatal(err)
	}

	if err := client.Get("/path", nil); err == nil {
		t.Fatal("Error was nil!")
	}
}
