package main

import (
	"flag"
	"fmt"
	"github.com/zpatrick/rclient"
	"log"
)

type Repository struct {
	Name string `json:"name,omitempty"`
}

func main() {
	username := flag.String("u", "", "username for your github account")
	password := flag.String("p", "", "password for your github account")
	flag.Parse()

	if *username == "" || *password == "" {
		log.Fatal("username and password are required")
	}

	client, err := rclient.NewRestClient("https://api.github.com")
	if err != nil {
		log.Fatal(err)
	}

	var repo Repository
	request := Repository{Name: "my_sample_repo"}

	// first, attempt to create a repo without auth - this will give us a 401
	if err := client.Post("/user/repos", request, &repo); err != nil {
		fmt.Printf("Failed to create repository without adding authentication: %v\n", err)
	}

	// next, try again but add auth to the request
	if err := client.Post("/user/repos", request, &repo, rclient.BasicAuth(*username, *password)); err != nil {
		log.Fatalf("Failed to create repository: %v", err)
	}

	fmt.Printf("Successfully created repository %s\n", repo.Name)

	// or, set basic auth with every request the client makes
	client, err = rclient.NewRestClient("https://api.github.com", rclient.RequestOptions(rclient.BasicAuth(*username, *password)))
	if err != nil {
		log.Fatal(err)
	}

	path := fmt.Sprintf("/repos/%s/%s", *username, repo.Name)
	if err := client.Delete(path, nil, nil); err != nil {
		log.Fatalf("Failed to delete repository: %v", err)
	}

	fmt.Printf("Successfully deleted repository %s\n", repo.Name)
}
