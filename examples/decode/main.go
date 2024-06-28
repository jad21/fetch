package main

import (
	"fmt"
	"log"

	"github.com/jad21/fetch"
)

const url = "https://api.github.com/users/%s"

type GitHubUser struct {
	Name     string `json:"name"`
	Company  string `json:"company"`
	Location string `json:"location"`
}

func main() {
	USERNAME := "jad21"

	rsp, err := fetch.Get(fmt.Sprintf(url, USERNAME), nil)
	if err != nil {
		log.Fatalf("could not fetch [%s] because: %s", url, err)
	}

	var user GitHubUser
	if err := rsp.Decode(&user); err != nil {
		log.Fatalf("could not fetch [%s] because: %s", url, err)
	}

	fmt.Printf("Name: %s\nCompany: %s\nLocation: %s\n", user.Name, user.Company, user.Location)
	/*
		Output:
		Name: Jose Angel
		Company:
		Location: Merida - Venezuela
	*/
}
