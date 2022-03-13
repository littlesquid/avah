package usecase

import (
	dataservice "avah/dataservice"
	"avah/model"
)

func ExecuteRegister(sheetId string, source model.Source) {

	dataRepository := dataservice.DBConfig{}
	dataRepository.InitDbConfig()
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
