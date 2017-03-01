package rclient

import (
	//"errors"
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
	t.Skip("todo")
}

func TestClientOptions(t *testing.T) {
	t.Skip("todo")
}

func TestRequestOptions(t *testing.T) {
	t.Skip("todo")
}

/*
func TestClientSenderError(t *testing.T) {
	doer := DoerFunc(func(req *http.Request) (*http.Response, error) {
		return nil, errors.New("some error")
	})

	client, err := NewRestClient("", WithRequestDoer(doer))
	if err != nil {
		t.Fatal(err)
	}

	if err := client.Get("/path", nil); err == nil {
		t.Fatal("Error was nil!")
	}
}

func TestClientReaderError(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		write(t, w, 200, person{"John Doe", 30})
	}

	reader := func(resp *http.Response, v interface{}) error {
		return errors.New("some error")
	}

	client, server := newClientAndServer(t, handler, WithReader(reader))
	defer server.Close()

	var p person
	if err := client.Get("/path", &p); err == nil {
		t.Fatal("Error was nil!")
	}
}
*/
