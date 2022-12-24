package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func main() {
	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile("/home/brasey/.gcp/shopgoodwill-scraper-3ef981e163aa.json")
	client, err := firestore.NewClient(ctx, "shopgoodwill-scraper", sa)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	search := client.Doc("config/search")

	m := map[string][]string{
		"terms": {
			"coach",
		},
	}

	_, err = search.Set(ctx, m)
	if err != nil {
		panic(err)
	}

	doc, err := search.Get(ctx)
	if err != nil {
		panic(err)
	}

	var searchRead map[string][]string

	if err = doc.DataTo(&searchRead); err != nil {
		panic(err)
	}

	for _, term := range searchRead["terms"] {
		fmt.Println(term)
	}
}
