package models

type AllowedAppLanguages struct {
	AllowAppLanguagesId int    `json:"-" gorm:"primary_key"`
	BoxOfficeId         int    `json:"-"`
	LanguageCode        string `json:"languageCode"`
	DisplayOrder        int    `json:"displayOrder"`
}
