package respository

import "crm-backend/internal/models"

type UserDaoInterface interface {
	CreateUser(user *models.User) error
	FindUserById(id string) *models.User
	UpdateUser(user *models.User)
	DeleteUserById(id string)
}
