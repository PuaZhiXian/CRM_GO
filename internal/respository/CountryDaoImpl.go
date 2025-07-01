package respository

import (
	"crm-backend/internal/models"
	"gorm.io/gorm"
)

type CountryDao struct {
	DB *gorm.DB
}

func (u *CountryDao) GetCountryByRiskLevel(riskLevels []string) ([]models.Country, error) {
	var countries []models.Country
	if len(riskLevels) == 0 {
		return countries, nil
	}
	err := u.DB.Where("RISK_LEVEL IN ?", riskLevels).Find(&countries).Error
	return countries, err
}
