package validator

import (
	api "crm-backend/gen"
	"crm-backend/internal/models"
	"crm-backend/internal/respository"
	"slices"
	"strings"
	"sync"
)

type UserValidator struct {
	CountryDao respository.CountryDaoInterface
}

func (u *UserValidator) ValidateUsers(users []models.User, dbUserChannel <-chan models.User, respWrapper *api.CreateUserRespWrapper) (validRecord []models.User) {
	var wg sync.WaitGroup
	invalidNationalityNResidentialSet := make(map[string]struct{})
	invalidDuplicateSet := make(map[string]struct{})

	wg.Add(2)
	go func() {
		defer wg.Done()
		for _, data := range u.validateNationalityNResidential(users) {
			strings.TrimSpace(data.Name + "_" + data.Nationality + "_" + data.Residential)
			invalidNationalityNResidentialSet[createUserKey(data)] = struct{}{}
		}
	}()
	go func() {
		defer wg.Done()
		for _, data := range validateDuplicateEntity(users, dbUserChannel) {
			invalidDuplicateSet[createUserKey(data)] = struct{}{}
		}
	}()
	wg.Wait()

	for _, data := range users {
		key := createUserKey(data)
		errMsg := ""
		if _, exists := invalidNationalityNResidentialSet[key]; exists {
			errMsg += "invalid nationality or residential "
		}

		if _, exists := invalidDuplicateSet[key]; exists {
			errMsg += "duplicate entity "
		}

		if len(errMsg) > 0 {
			*respWrapper.FailedUser = append(*respWrapper.FailedUser, api.CreateUserRespFailed{
				Name:        &data.Name,
				Nationality: &data.Name,
				Residential: &data.Residential,
				Age:         &data.Age,
				Reason:      &errMsg,
			})
		} else {
			validRecord = append(validRecord, data)
		}
	}

	return validRecord
}

func (u *UserValidator) validateNationalityNResidential(users []models.User) (invalidRecord []models.User) {
	countries, _ := u.CountryDao.GetCountryByRiskLevel([]string{"LOW"})
	allowNationality, allowResidential := prepareNationalityAndResidentialSet(countries)

	invalidRecord = make([]models.User, 0, 10)
	var mu sync.Mutex
	var wg sync.WaitGroup
	for _, user := range users {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !(slices.Contains(allowNationality, user.Nationality) && slices.Contains(allowResidential, user.Residential)) {
				mu.Lock()
				invalidRecord = append(invalidRecord, user)
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	return invalidRecord
}

func validateDuplicateEntity(users []models.User, dbUserChannel <-chan models.User) (invalidRecord []models.User) {
	invalidRecord = make([]models.User, 0, len(users))
	userMap := make(map[string]models.User)
	for i := 0; i < len(users); i++ {
		userMap[createUserKey(users[i])] = users[i]
	}

	for dbUser := range dbUserChannel {
		if pointer, exist := userMap[createUserKey(dbUser)]; exist {
			invalidRecord = append(invalidRecord, pointer)
		}
	}

	return invalidRecord
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
	return len(user.Nationality) > 0 && len(user.Residential) > 0 && slices.Contains(allowNationality, user.Nationality) && slices.Contains(allowResidential, user.Residential)
}

func createUserKey(user models.User) string {
	return strings.TrimSpace(user.Name + "_" + user.Nationality + "_" + user.Residential)
}
