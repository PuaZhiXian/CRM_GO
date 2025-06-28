package models

type Country struct {
	CountryCode uint   `json:"countryCode" gorm:"primaryKey;autoIncrement;column:COUNTRY_CODE" `
	CountryName string `json:"countryName" gorm:"size:100;not null;column:COUNTRY_NAME"`
	ISO         string `json:"iso" gorm:"size:2;not null;unique;column:ISO"`
	Nationality string `json:"nationality" gorm:"size:100;not null;column:NATIONALITY"`
	RiskLevel   string `json:"riskLevel" gorm:"size:100;column:RISK_LEVEL" `
}

func (Country) TableName() string {
	return "country"
}
