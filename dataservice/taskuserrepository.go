package dataservice

import (
	"avah/model"
	"fmt"
)

type TaskUserRepository struct {
	DBConfig *DBConfig
}

func (repository *TaskUserRepository) Insert(taskUserModel *model.TaskUser) {
	fmt.Println("insert task user profile")
	_, err := repository.DBConfig.DB.Model(taskUserModel).Insert()
	if err != nil {
		fmt.Errorf("insert task failed: ", err)
	}
}
