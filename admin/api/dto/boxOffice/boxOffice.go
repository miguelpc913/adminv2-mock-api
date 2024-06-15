package dtoBoxOffice

import "github.com/tiqueteo/adminv2-mock-api/db/models"

type AllowedLanguagesDTO struct {
	LanguageCode string `json:"languageCode"`
	DisplayOrder int    `json:"displayOrder"`
}

type BoxOfficePost struct {
	Status             models.BoxOfficeStatus `json:"status"`
	Alias              string                 `json:"alias"`
	SalesGroupId       int                    `json:"salesGroupId"`
	BillPrefix         string                 `json:"billPrefix"`
	CashCountThreshold int                    `json:"cashCountThreshold"`
	CashCountAttempts  int                    `json:"cashCountAttempts"`
}

type BoxOfficeBasicConfigurations struct {
	Status         models.BoxOfficeStatus `json:"status"`
	Alias          string                 `json:"alias"`
	SalesGroupId   int                    `json:"salesGroupId"`
	OfflineEnabled bool                   `json:"offlineEnabled"`
}

type BoxOfficeCashCount struct {
	CashCountThreshold int `json:"cashCountThreshold"`
	CashCountAttempts  int `json:"cashCountAttempts"`
}

type BoxOfficePresentations struct {
	AllowSelectMainProduct    bool `json:"allowSelectMainProduct"`
	GroupProductsByCategory   bool `json:"groupProductsByCategory"`
	ShowOnlyAvailableEvents   bool `json:"showOnlyAvailableEvents"`
	ShowEventFinalHour        bool `json:"showEventFinalHour"`
	OfferNextAvailableSession bool `json:"offerNextAvailableSession"`
	ShowMaxAvailability       bool `json:"showMaxAvailability"`
}

type BoxOfficeFunctionalities struct {
	AllowPromotionalCodes      bool `json:"allowPromotionalCodes"`
	SellerSelectorEnabled      bool `json:"sellerSelectorEnabled"`
	CustomerSelectorEnabled    bool `json:"customerSelectorEnabled"`
	PurchaseFormIntervalNeeded int  `json:"purchaseFormIntervalNeeded"`
}

type BoxOfficeLanguages struct {
	AllowedTicketLanguages []AllowedLanguagesDTO `json:"allowedTicketLanguages"`
	AllowedAppLanguages    []AllowedLanguagesDTO `json:"allowedAppLanguages"`
	LanguageCode           string                `json:"languageCode"`
}

type BoxOfficePrintSettings struct {
	PrintTicket                  bool     `json:"printTicket"`
	OptionalPrintTicket          bool     `json:"optionalPrintTicket"`
	PrintTicketPrice             bool     `json:"printTicketPrice"`
	OptionalPrintTicketPrice     bool     `json:"optionalPrintTicketPrice"`
	PrintSummary                 bool     `json:"printSummary"`
	OptionalPrintSummary         bool     `json:"optionalPrintSummary"`
	HighlightPrintedReservations bool     `json:"highlightPrintedReservations"`
	SingleDocPrint               bool     `json:"singleDocPrint"`
	AllowedTicketGroupTypes      []string `json:"allowedTicketGroupTypes"`
	PrintCashCount               bool     `json:"printCashCount"`
}

type BoxOfficePaymentRequest struct {
	ShowChangeCalculator      int  `json:"showChangeCalculator"`
	AskCashForChange          bool `json:"askCashForChange"`
	Calculator                bool `json:"calculator"`
	ShowDeferredPaymentMethod bool `json:"showDeferredPaymentMethod"`
}

type BoxOfficeValidations struct {
	ValidationMethod int `json:"validationMethod"`
	VerifierId       int `json:"verifierId"`
}

type BoxOfficeAdvancedSettings struct {
	AllowConfigPrinter         bool `json:"allowConfigPrinter"`
	AllowLoginHistory          bool `json:"allowLoginHistory"`
	AllowChangePassword        bool `json:"allowChangePassword"`
	AllowSellAndEditPastEvents bool `json:"allowSellAndEditPastEvents"`
	PrintMultiple              bool `json:"printMultiple"`
	CloseDetailsPrintTickets   bool `json:"closeDetailsPrintTickets"`
	CloseDetailsPrintReceipt   bool `json:"closeDetailsPrintReceipt"`
	AvailabilityThreshold      int  `json:"availabilityThreshold"`
}
