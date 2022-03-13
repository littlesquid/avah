package dataservice

import (
	"avah/model"
	"fmt"
)

type UserProfileRepository struct {
	DBConfig *DBConfig
}

func (repository *UserProfileRepository) Insert(userProfile *model.UserProfile) {
	fmt.Println("insert user profile")
	_, err := repository.DBConfig.DB.Model(userProfile).Insert()
	if err != nil {
		fmt.Errorf("insert user profile failed: ", err)
	}
}
