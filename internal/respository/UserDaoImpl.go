package respository

import (
	"crm-backend/internal/models"
	"crm-backend/pkg/util"
	"gorm.io/gorm"
	"time"
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

func (u *UserDao) GetDataCount() (int64, error) {
	var total int64
	if err := u.DB.Model(&models.User{}).Count(&total).Error; err != nil {
		return total, err
	}
	return total, nil
}

func (u *UserDao) FindUserPage(size int, page int, column []string) ([]models.User, error) {
	var chunk []models.User

	query := u.DB
	if len(column) > 0 {
		query = query.Select(column)
	}

	err := query.Limit(size).Offset(page * size).Find(&chunk).Error
	if err != nil {
		return nil, err
	}

	return chunk, nil
}
