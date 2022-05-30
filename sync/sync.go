package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	c := http.Client{}

	resp, err := c.Get("http://google.com")

	if err != nil {
		log.Println("Error", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}

	fmt.Printf("Body : %s", body)

	// youtubeMain(&c)
}
