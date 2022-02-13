package service

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type GoogleSheetService struct {
	srv *sheets.Service
}

func (sheetService *GoogleSheetService) ConnectGoogleSheet(client *http.Client) {
	ctx := context.Background()

	// Creating an Interface
	rows := make([][]interface{}, 0)
	// Adding the first line of the sheet
	rows = append(rows, []interface{}{"Last Updated", "Title", "Article Link"})

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	sheetService.srv = srv
}

func (sheetService *GoogleSheetService) ReadGoogleSheet() interface{} {
	spreadsheetId := viper.GetString("sheet.id")
	rangeData := "sheet1!A1:E"
	resp, err := sheetService.srv.Spreadsheets.Values.Get(spreadsheetId, rangeData).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("data not found")
		return nil
	}

	return resp.Values
}
