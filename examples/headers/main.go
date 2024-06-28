package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jad21/fetch"
)

func main() {
	headers := http.Header{
		"Content-Type": []string{"application/json"},
		"User-Agent":   []string{"My-User-Agent"},
	}

	f := fetch.New(fetch.WithHeader(headers))
	rsp, err := f.Get("https://httpbin.org/headers", nil)
	if err != nil {
		log.Fatalf("could not fetch data from target because: %s", err)
	}
	fmt.Println(rsp.String())
	/*
		...
			"User-Agent": "My-User-Agent",
		...
	*/
}
