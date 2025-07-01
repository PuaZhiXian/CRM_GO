package respository

import "crm-backend/internal/models"

type UserDaoInterface interface {
	CreateUser(user *models.User) error
	FindUserById(id string) *models.User
	UpdateUser(user *models.User)
	DeleteUserById(id string)
	GetDataCount() (int64, error)
	FindUserPage(size int, page int, column []string) ([]models.User, error)
}
