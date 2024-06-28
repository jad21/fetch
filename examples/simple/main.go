package main

import (
	"log"

	"github.com/jad21/fetch"
)

const url = "https://api.github.com/users/jad21"

func main() {
	rsp, err := fetch.Get(url, nil)
	if err != nil {
		log.Fatalf("could not fetch [%s] because: %s", url, err)
	}

	body, err := rsp.ToString()
	if err != nil {
		log.Fatalf("could not retrieve body because: %s", err)
	}

	log.Println(body)
}
