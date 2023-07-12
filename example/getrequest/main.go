package main

import (
	"calista/httpclient"
	"fmt"
	"log"
)

func main() {
	httpCli, err := httpclient.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	var resp struct {
		ID          int64   `json:"id"`
		Title       string  `json:"title"`
		Price       float64 `json:"price"`
		Description string  `json:"description"`
		Category    string  `json:"category"`
		Image       string  `json:"image"`
		Rating      struct {
			Rate  float64 `json:"rate"`
			Count int64   `json:"count"`
		} `json:"rating"`
	}
	err = httpCli.Get("https://fakestoreapi.com/products/1", nil, &resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
