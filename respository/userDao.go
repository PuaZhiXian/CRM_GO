package respository

import (
	"crm-backend/models"
	"database/sql"
	"log"
	"time"
)

type UserDao struct {
	DB *sql.DB
}

func (u *UserDao) CreateUser(user *models.User) {
	log.Println("inserting data to db ", user)
	//TODO INSERT DATA TO DB
}
func (u *UserDao) FindUserById(id string) *models.User {
	//TODO RETURN USER
	user := models.User{
		Id:          "1",
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
