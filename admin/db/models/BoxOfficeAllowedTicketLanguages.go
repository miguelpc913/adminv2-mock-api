package models

type AllowedTicketLanguages struct {
	AllowedTicketLanguagesId int    `json:"-" gorm:"primary_key"`
	BoxOfficeId              int    `json:"-"`
	LanguageCode             string `json:"languageCode"`
	DisplayOrder             int    `json:"displayOrder"`
}
