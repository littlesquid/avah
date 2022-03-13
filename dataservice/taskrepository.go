package dataservice

import (
	"avah/model"
	"fmt"
)

type TaskRepository struct {
	DBConfig *DBConfig
}

func (repository *TaskRepository) Insert(taskModel *model.Task) {
	fmt.Println("insert task")
	_, err := repository.DBConfig.DB.Model(taskModel).Insert()
	if err != nil {
		fmt.Errorf("insert task failed: ", err)
	}
}
