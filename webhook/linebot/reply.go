package webhook

import (
	"avah/oauth"
	"avah/security"
	service "avah/service/googleapi"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

func Reply(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is reply function!")

	requestBody, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	lineWebHookRequest := LineWebHookRequest{}
	if err := json.Unmarshal(requestBody, &lineWebHookRequest); err != nil {
		panic(err)
	}

	fmt.Printf("got request: %s \n", string(requestBody[:]))

	googleApiClient := oauth.GoogleApiLogin()

	googleApiService := service.GoogleSheetService{}
	googleApiService.ConnectGoogleSheet(googleApiClient)

	spreadsheetId := viper.GetString("sheet.id")
	sheetName := "sheet1!A"
	readRange := fmt.Sprintf("%v1:E", sheetName)

	sheetData := googleApiService.ReadGoogleSheet(spreadsheetId, readRange)

	aesEncryption := security.AesEncryption{}
	aesEncryption.GenerateKey()
	aesEncryption.Init()

	userId := lineWebHookRequest.Events[0].Source.UserID

	encryptedUserId := aesEncryption.Encrypt(userId)

	rangeData := fmt.Sprintf("%s%v", sheetName, len(sheetData.Values)+1)

	rows := make([][]interface{}, 0)
	rows = append(rows, []interface{}{encryptedUserId})

	googleApiService.WriteGoogleSheet(spreadsheetId, rangeData, rows)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v", sheetData)))
}
