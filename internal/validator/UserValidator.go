package validator

import (
	api "crm-backend/gen"
	"crm-backend/internal/models"
	"crm-backend/internal/respository"
	"slices"
	"sync"
)

type UserValidator struct {
	CountryDao respository.CountryDaoInterface
}

func (u *UserValidator) ValidateUsers(users []models.User, respWrapper *api.CreateUserRespWrapper) (validRecord []models.User) {
	countries, _ := u.CountryDao.GetCountryByRiskLevel([]string{"LOW"})
	allowNationality, allowResidential := prepareNationalityAndResidentialSet(countries)

	validRecord = make([]models.User, 0, 10)
	var mu sync.Mutex
	var wg sync.WaitGroup
	for _, user := range users {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if validateUser(&user, allowNationality, allowResidential) {
				mu.Lock()
				validRecord = append(validRecord, user)
				mu.Unlock()
			} else {
				mu.Lock()
				*respWrapper.FailedUser = append(*respWrapper.FailedUser, api.CreateUserRespFailed{
					Name:        &user.Name,
					Nationality: &user.Name,
					Residential: &user.Residential,
					Age:         &user.Age,
					Reason:      func(s string) *string { return &s }("Just Invalid no ask"),
				})
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	return validRecord
}

func prepareNationalityAndResidentialSet(countries []models.Country) ([]string, []string) {
	allowNationality := make([]string, 0, len(countries))
	allowResidential := make([]string, 0, len(countries))

	for _, country := range countries {
		allowNationality = append(allowNationality, country.Nationality)
		allowResidential = append(allowResidential, country.CountryName)
	}

	return allowNationality, allowResidential
}

func validateUser(user *models.User, allowNationality []string, allowResidential []string) bool {
	return slices.Contains(allowNationality, user.Nationality) && slices.Contains(allowResidential, user.Residential)
}
