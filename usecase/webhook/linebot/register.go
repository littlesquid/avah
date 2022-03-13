package usecase

import (
	adapter "avah/adapter/googleclient"
	dataservice "avah/dataservice"
	"avah/model"
	"fmt"

	"google.golang.org/api/sheets/v4"
)

func ExecuteRegister(sheetId string, source model.Source) {

	dataRepository := dataservice.DBConfig{
		UserName: "root",
		Password: "password",
		Host:     "localhost",
		Port:     "5432",
		DbName:   "avah",
		Timeout:  20,
		SSLMode:  "disabled",
	}

	dataRepository.OpenConnection()

	defer dataRepository.DB.Close()

	taskRepository := dataservice.TaskRepository{
		DBConfig: &dataRepository,
	}
	userProfileRepository := dataservice.UserProfileRepository{
		DBConfig: &dataRepository,
	}
	taskUserRepository := dataservice.TaskUserRepository{
		DBConfig: &dataRepository,
	}

	task := model.Task{
		SourceId: source.UserID,
		SheetId:  sheetId,
		IsActive: 1,
	}

	taskRepository.Insert(&task)

	userProfileRepository.Insert(&model.UserProfile{
		UserId:   source.UserID,
		UserType: source.Type,
	})

	taskUserRepository.Insert(&model.TaskUser{
		TaskId: task.Id,
		UserId: source.UserID,
	})
}

func retreiveDataFromSheet(sheetId string) *sheets.ValueRange {
	googleApiClient := adapter.GoogleApiLogin()

	googleApiService := adapter.GoogleSheetService{}
	err := googleApiService.ConnectGoogleSheet(googleApiClient)
	if err != nil {
		fmt.Errorf("connect google sheet failed: ", err)
	}

	sheetName := "Data!A"
	readRange := fmt.Sprintf("%v1:E", sheetName)

	sheetData := googleApiService.ReadGoogleSheet(sheetId, readRange)

	fmt.Printf("rangeData: %s \n", sheetData.Values[0][0])

	return sheetData
}
