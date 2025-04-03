package main

import (
	"context"
	"errors"
	"github.com/jomei/notionapi"
	"time"
)

type PageDB struct {
	title       string
	date        int64
	category    string
	description string
	amount      int
}

func queryDB(client *notionapi.Client, dbID string) ([]PageDB, error) {
	var query notionapi.DatabaseQueryRequest
	dbQuery, err := client.Database.Query(context.Background(), notionapi.DatabaseID(dbID), &query)
	if err != nil {
		return nil, errors.New("Error when querying the database: " + err.Error())
	}

	var allRows []PageDB
	for _, pages := range dbQuery.Results {

		pageDB := PageDB{}
		for _, page := range pages.Properties {
			switch x := page.(type) {
			case *notionapi.RichTextProperty:
				for _, richText := range x.RichText {
					pageDB.description = string(richText.Text.Content)
				}
			case *notionapi.TitleProperty:
				for _, title := range x.Title {
					pageDB.title = string(title.Text.Content)
				}
			case *notionapi.NumberProperty:
				pageDB.amount = int(x.Number)
			case *notionapi.DateProperty:
				parsedDate, err := time.Parse(time.RFC3339, x.Date.Start.String())
				if err != nil {
					errors.New("Error when parsing date: " + err.Error())
				}
				pageDB.date = parsedDate.Unix()
			case *notionapi.MultiSelectProperty:
				if x.MultiSelect[0].ID == "KPxs" || x.MultiSelect[0].ID == ":w\\r" {
					pageDB.category = x.MultiSelect[0].Name
				}

			}

		}
		allRows = append(allRows, pageDB)
	}
	return allRows, nil
}
