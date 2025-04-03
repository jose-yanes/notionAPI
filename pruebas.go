package main

import (
	"context"
	"fmt"
	"github.com/jomei/notionapi"
)

func testing(client *notionapi.Client, dbID string) {
	dbQuery, err := client.Block.GetChildren(context.Background(), notionapi.BlockID(dbID), nil)
	if err != nil {
		fmt.Println(err)
	}
	for _, block := range dbQuery.Results {
		fmt.Println(block)
	}
}
