package main

import (
	"log"

	"github.com/badoux/goscraper"
	"github.com/emwalker/digraph/server/api"
	"github.com/emwalker/digraph/server/dgraph"
	"github.com/labstack/echo"
)

func titleFetcher(url string) (string, error) {
	log.Println("fetching page", url, "to get the title ...")
	page, err := goscraper.Scrape(url, 5)
	if err != nil {
		return "", err
	}
	return page.Preview.Title, nil
}

func main() {
	s := &api.Server{
		Engine: echo.New(),
		Schema: dgraph.LoadSchema("./schema.graphql"),
	}

	log.Fatal(s.ListenAndServe())
}
