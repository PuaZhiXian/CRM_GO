package service

import (
	api "crm-backend/gen"
	"crm-backend/internal/models"
	"crm-backend/internal/respository"
	"crm-backend/internal/validator"
	"log"
	"sync"
)

type UserService struct {
	UserDao    respository.UserDaoInterface
	CountryDao respository.CountryDaoInterface
}

func (s *UserService) CreateUser(user *models.User) (string, error) {
	log.Println("[CreateUser] Start")
	if err := s.UserDao.CreateUser(user); err != nil {
		return "", err
	}
	return "User Created", nil
}

func (s *UserService) CreateUserByBulk(users *[]models.User) api.CreateUserRespWrapper {
	log.Println("[CreateUserByBulk] Start")

	success := make([]api.CreateUserRespSuccess, 0, 10)
	failed := make([]api.CreateUserRespFailed, 0, 10)
	respWrapper := &api.CreateUserRespWrapper{
		SuccessUser: &success,
		FailedUser:  &failed,
	}

	validator := validator.UserValidator{
		CountryDao: s.CountryDao,
	}
	validRecord := validator.ValidateUsers(*users, respWrapper)
	s.insertUserToDB(validRecord, respWrapper)
	return *respWrapper
}

func (s *UserService) insertUserToDB(validRecord []models.User, respWrapper *api.CreateUserRespWrapper) {
	var mu sync.Mutex
	var wg sync.WaitGroup
	for _, validUser := range validRecord {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := s.UserDao.CreateUser(&validUser); err != nil {
				mu.Lock()
				*respWrapper.FailedUser = append(*respWrapper.FailedUser, api.CreateUserRespFailed{
					Name:        &validUser.Name,
					Nationality: &validUser.Name,
					Residential: &validUser.Residential,
					Age:         &validUser.Age,
					Reason:      func(s string) *string { return &s }("Internal Err"),
				})
				mu.Unlock()
			} else {
				mu.Lock()
				*respWrapper.SuccessUser = append(*respWrapper.SuccessUser, api.CreateUserRespSuccess{
					Name:        &validUser.Name,
					Nationality: &validUser.Name,
					Residential: &validUser.Residential,
					Age:         &validUser.Age,
				})
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
}
