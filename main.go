package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/jomei/notionapi"
	"os"
)

func main() {

	// Get env variables
	godotenv.Load()
	token := os.Getenv("TOKEN")
	dbID := os.Getenv("DB_ID")
	dbIncome := os.Getenv("DB_INCOME")
	// Creates notionApi client with Token
	client := notionapi.NewClient(notionapi.Token(token))

	// queryDB queries and formats the results on the DB
	// Returns []PageDB
	dbResults, _ := queryDB(client, dbID)

	necessityAmt := 0
	wantAmt := 0
	savingAmt := 0

	for _, dbResult := range dbResults {
		switch dbResult.category {
		case "Gasto Necesario":
			necessityAmt += dbResult.amount
		case "Gasto Ocio":
			wantAmt += dbResult.amount
		case "Ahorro / Inversion":
			savingAmt += dbResult.amount
		}
	}

	fmt.Printf("Necessity amount: %d\nWants Amount: %d\nSavings Amount: %d\n", necessityAmt, wantAmt, savingAmt)

	testing(client, dbIncome)
}
