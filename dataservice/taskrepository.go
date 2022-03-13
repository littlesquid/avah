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

func (repository *TaskRepository) FindAll() []model.Task {
	fmt.Println("find all task")
	fmt.Println(repository.DBConfig)
	var tasks []model.Task
	err := repository.DBConfig.DB.Model(&tasks).Select()

	if err != nil {
		fmt.Println("insert task failed: ", err)
		return nil
	}

	return tasks
}
