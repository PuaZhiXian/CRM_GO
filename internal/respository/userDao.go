package respository

import (
	"crm-backend/internal/models"
	"crm-backend/pkg/util"
	"time"

	"gorm.io/gorm"
)

type UserDao struct {
	DB *gorm.DB
}

func (u *UserDao) CreateUser(user *models.User) error {
	result := u.DB.Create(user)
	if result.Error != nil {
		return util.ErrUnexpected
	}
	return nil
}

func (u *UserDao) FindUserById(id string) *models.User {
	user := models.User{
		Id:          1,
		Name:        "Hello World",
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	}
	return &user
}

func (u *UserDao) UpdateUser(user *models.User) {
	//TODO UPDATE USER
}
func (u *UserDao) DeleteUserById(id string) {
	//TODO DELETE USER
}
