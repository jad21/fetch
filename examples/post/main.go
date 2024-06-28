package main

import (
	"log"

	"github.com/jad21/fetch"
)

type Payload struct {
	Name  string
	Email string
}

func main() {
	payload := Payload{
		Name:  "Jose Delgado",
		Email: "esojangel@gmail.com",
	}

	rsp, err := fetch.IsJSON().Post("https://httpbin.org/post", fetch.NewReader(payload))
	if err != nil {
		log.Fatalf("could not login because: %s", err)
	}

	log.Println(rsp.String())
	/*
		...
		"json": {
			"Email": "esojangel@gmail.com",
			"Name": "Jose Delgado"
		},
		...
	*/
}
