package usecase

import (
	googleclient "avah/adapter/googleclient"
	lineclient "avah/adapter/lineclient"
	"avah/dataservice"
	"fmt"

	"google.golang.org/api/sheets/v4"
)

func ExecuteDailyReminder() {
	dataRepository := dataservice.DBConfig{}
	dataRepository.InitDbConfig()
	dataRepository.OpenConnection()

	defer dataRepository.DB.Close()

	taskRepository := dataservice.TaskRepository{
		DBConfig: &dataRepository,
	}

	tashUserRepository := dataservice.TaskUserRepository{
		DBConfig: &dataRepository,
	}

	tasks := taskRepository.FindAll()

	for _, task := range tasks {
		fmt.Println(task)
		sheetData := retreiveDataFromSheet(task.SheetId)

		message := fmt.Sprintf("%s", sheetData.Values[0][0])

		taskUser := tashUserRepository.FindByTaskId(task.Id)

		lineclient.Remind(message, taskUser[0].UserId)
	}
}

func retreiveDataFromSheet(sheetId string) *sheets.ValueRange {
	googleApiClient := googleclient.GoogleApiLogin()

	googleApiService := googleclient.GoogleSheetService{}
	err := googleApiService.ConnectGoogleSheet(googleApiClient)
	if err != nil {
		fmt.Errorf("connect google sheet failed: ", err)
	}

	sheetName := "Data!A"
	readRange := fmt.Sprintf("%v1:E", sheetName)

	sheetData := googleApiService.ReadGoogleSheet(sheetId, readRange)

	fmt.Println("rangeData: %s \n", sheetData.Values[0][0])

	return sheetData
}
