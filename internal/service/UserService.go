package service

import (
	"crm-backend/internal/models"
	"crm-backend/internal/respository"
	"log"
)

type UserService struct {
	UserDao respository.UserDaoInterface
}

func (s *UserService) CreateUser(user *models.User) (error, string) {
	log.Println("[CreateUser] Start")
	if err := s.UserDao.CreateUser(user); err != nil {
		return err, ""
	}
	return nil, "User Created"
}
