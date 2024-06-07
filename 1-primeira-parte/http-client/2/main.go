package main

import (
	"io"
	"net/http"
)

func main() {
	c := http.Client{}
	req, err := http.NewRequest("GET", "https://google.com", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Accept", "application/json")
	rest, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer rest.Body.Close()

	body, err := io.ReadAll(rest.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
