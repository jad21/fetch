package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jad21/fetch"
)

func main() {
	cookie := http.Cookie{
		Name:  "cookie1",
		Value: "value1",
	}

	f := fetch.New(fetch.WithCookies(cookie))
	rsp, err := f.Get("https://httpbin.org/cookies", nil)
	if err != nil {
		log.Fatalf("could not fetch data from target because: %s", err)
	}
	fmt.Println(rsp.String())
	/*
		...
			"cookies": {
				"cookie1": "value1"
			}
		...
	*/
}
