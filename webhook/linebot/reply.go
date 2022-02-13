package webhook

import (
	"avah/oauth"
	service "avah/service/googleapi"
	"fmt"
	"io/ioutil"
	"net/http"
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

	fmt.Printf("got request: %s \n", string(requestBody[:]))

	googleApiClient := oauth.GoogleApiLogin()

	googleApiService := service.GoogleSheetService{}
	googleApiService.ConnectGoogleSheet(googleApiClient)

	sheetData := googleApiService.ReadGoogleSheet()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v", sheetData)))
}
