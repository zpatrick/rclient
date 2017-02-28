# RClient

[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/zpatrick/rclient/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/zpatrick/rclient)](https://goreportcard.com/report/github.com/zpatrick/rclient)
[![Go Doc](https://godoc.org/github.com/zpatrick/rclient?status.svg)](https://godoc.org/github.com/zpatrick/rclient)

## Overview
RClient is a Go package for interacting with REST APIs. 
It aims to have a clean, readable API for simple use cases, while having the flexibility to accommodate complex application requirements.

## Getting Started
The following snippet shows RClient interacting with Github's API: 
```
package main

import (
        "github.com/zpatrick/rclient"
        "log"
)

type Repository struct {
        Name        string `json:"name"`
}

func main() {
        client, err := rclient.NewRestClient("https://api.github.com")
        if err != nil {
                log.Fatal(err)
        }

        var repos []Repository
        if err := client.Get("/users/zpatrick/repos", &repos); err != nil {
                log.Fatal(err)
        }

        log.Println(repos)
}
```

# License
This work is published under the MIT license.

Please see the `LICENSE` file for details.
