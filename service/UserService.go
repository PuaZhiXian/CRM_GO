package service

import (
	"crm-backend/models"
	"crm-backend/respository"
	"log"
	"os"
)

type UserService struct {
	UserDao respository.UserDaoInterface
}

func (s *UserService) CreateUser(user *models.User) string {
	logger := log.New(os.Stdout, "[CreateUser] ", log.LstdFlags)
	logger.SetFlags(0)
	logger.Println("Start")
	
	s.UserDao.CreateUser(user)
	return "User Created"
}
