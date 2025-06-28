package respository

import "crm-backend/internal/models"

type CountryDaoInterface interface {
	GetCountryByRiskLevel(riskLevel []string) ([]models.Country, error)
}
