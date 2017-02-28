package main

import (
	"flag"
	"github.com/zpatrick/rclient"
	"log"
	"fmt"
)

type Repository struct {
	Name        string `json:"name"`
}

func main() {
	username := flag.String("u", "zpatrick", "username for your github account")
	flag.Parse()

	client, err := rclient.NewRestClient("https://api.github.com")
	if err != nil {
		log.Fatal(err)
	}
	
	var repos []Repository
	path := fmt.Sprintf("/users/%s/repos", *username)
        if err := client.Get(path, &repos); err != nil {
                log.Fatal(err)
        }

	fmt.Printf("Repos for %s: \n", *username)
        for _, r := range repos {
                fmt.Println(r.Name)
        }
}

