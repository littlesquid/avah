package handler

import (
	adapter "avah/adapter/googleclient"
	"avah/security"
	webhook "avah/usecase/webhook/linebot"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

func HandleRequests() {
	http.HandleFunc("/reply", webhook.Reply)
	http.HandleFunc("/callback", webhook.Callback)
	http.HandleFunc("/health", healthCheck)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {

	googleApiClient := adapter.GoogleApiLogin()

	googleApiService := adapter.GoogleSheetService{}
	googleApiService.ConnectGoogleSheet(googleApiClient)

	spreadsheetId := viper.GetString("sheet.id")
	rangeData := "sheet1!A1:E"

	sheetData := googleApiService.ReadGoogleSheet(spreadsheetId, rangeData)

	aesEncryption := security.AesEncryption{}
	aesEncryption.LoadKey()
	aesEncryption.Init()

	resp := ""
	for _, row := range sheetData.Values {
		// Print columns A and E, which correspond to indices 0 and 4.
		value := fmt.Sprintf("%v", row[0])
		resp = fmt.Sprintf("%s\n", aesEncryption.Decrypt(value))
	}
	//fmt.Printf("decrypted userId: %s \n", aesEncryption.Decrypt(encryptedUserId))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v", resp)))
	return
}
