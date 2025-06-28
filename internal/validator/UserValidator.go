package validator

import (
	api "crm-backend/gen"
	"crm-backend/internal/models"
	"fmt"
	"slices"
	"sync"
)

var SUPPORTED_NATIONALTIY = []string{"Malaysia", "Singapore"}
var SUPPORTED_RESIDENTIAL = []string{"Malaysian", "Singaporen"}

func ValidateUsers(users []models.User, respWrapper *api.CreateUserRespWrapper) (validRecord []models.User) {
	fmt.Println("ValidateUsers")

	validRecord = make([]models.User, 0, 10)
	var mu sync.Mutex
	var wg sync.WaitGroup
	for _, user := range users {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if validateUser(&user) {
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

func validateUser(user *models.User) bool {
	return slices.Contains(SUPPORTED_NATIONALTIY, user.Nationality) && slices.Contains(SUPPORTED_RESIDENTIAL, user.Residential)
}
