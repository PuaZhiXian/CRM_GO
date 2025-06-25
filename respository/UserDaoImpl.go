package respository

import (
	"crm-backend/models"
)

type UserDaoInterface interface {
	CreateUser(user *models.User)
	FindUserById(id string) *models.User
	UpdateUser(user *models.User)
	DeleteUserById(id string)
}
