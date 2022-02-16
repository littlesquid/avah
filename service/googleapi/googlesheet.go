package service

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type GoogleSheetService struct {
	srv *sheets.Service
}

func (sheetService *GoogleSheetService) ConnectGoogleSheet(client *http.Client) {
	ctx := context.Background()

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	sheetService.srv = srv
}

func (sheetService *GoogleSheetService) ReadGoogleSheet(sheetId string, rangeData string) *sheets.ValueRange {
	resp, err := sheetService.srv.Spreadsheets.Values.Get(sheetId, rangeData).Do()
	if err != nil {
		fmt.Printf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("data not found")
		return nil
	}

	return resp
}

func (sheetService *GoogleSheetService) WriteGoogleSheet(sheetId string, rangeData string, rows [][]interface{}) {

	rb := &sheets.BatchUpdateValuesRequest{
		ValueInputOption: "USER_ENTERED",
	}

	rb.Data = append(rb.Data, &sheets.ValueRange{
		Range:  rangeData,
		Values: rows,
	})

	_, err := sheetService.srv.Spreadsheets.Values.BatchUpdate(sheetId, rb).Do()
	if err != nil {
		fmt.Printf("Unable to write data to sheet: %v", err)
	}
}
