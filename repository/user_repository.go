package repository

import (
	"beckend/config"
	"beckend/model"
)

func FindUserByUsername(username string) (model.User, error) {
	var user model.User
	err := config.GetDB().First(&user, "username = ?", username).Error
	return user, err
}
func CreateUser(user model.User) (model.User, error) {
	err := config.GetDB().Create(&user).Error
	return user, err
}
