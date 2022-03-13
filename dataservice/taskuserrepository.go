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

func (repository *TaskUserRepository) FindByTaskId(taskId int64) []model.TaskUser {
	fmt.Println("find user by taskId")
	var taskUsers []model.TaskUser
	err := repository.DBConfig.DB.Model(&taskUsers).Where("task_user.task_id = ?", taskId).Select()

	if err != nil {
		fmt.Errorf("insert task failed: ", err)
		return nil
	}

	return taskUsers
}
