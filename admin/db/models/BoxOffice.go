package models

import (
	"time"

	"gorm.io/gorm"
)

type Language struct {
	LanguageCode string `json:"languageCode"`
	DisplayOrder int    `json:"displayOrder"`
}

type BoxOfficeStatus string

const (
	EnabledBoxOffice  BoxOfficeStatus = "enabled"
	DisabledBoxOffice BoxOfficeStatus = "disabled"
	PendingBoxOffice  BoxOfficeStatus = "deprecated"
)

type BoxOffice struct {
	BoxOfficeId                  int             `json:"boxOfficeId" gorm:"primary_key"`
	Status                       BoxOfficeStatus `json:"status"`
	Alias                        string          `json:"alias"`
	SalesGroupId                 int             `json:"salesGroupId"`
	BillPrefix                   string          `json:"billPrefix"`
	InstalledVersion             string          `json:"installedVersion"`
	OfflineEnabled               bool            `json:"offlineEnabled"`
	CashCountThreshold           int             `json:"cashCountThreshold"`
	CashCountAttempts            int             `json:"cashCountAttempts"`
	AllowSelectMainProduct       bool            `json:"allowSelectMainProduct"`
	GroupProductsByCategory      bool            `json:"groupProductsByCategory"`
	ShowOnlyAvailableEvents      bool            `json:"showOnlyAvailableEvents"`
	ShowEventFinalHour           bool            `json:"showEventFinalHour"`
	OfferNextAvailableSession    bool            `json:"offerNextAvailableSession"`
	ShowMaxAvailability          bool            `json:"showMaxAvailability"`
	AllowPromotionalCodes        bool            `json:"allowPromotionalCodes"`
	SellerSelectorEnabled        bool            `json:"sellerSelectorEnabled"`
	CustomerSelectorEnabled      bool            `json:"customerSelectorEnabled"`
	PurchaseFormIntervalNeeded   int             `json:"purchaseFormIntervalNeeded"`
	AllowedTicketLanguages       Language        `json:"allowedTicketLanguages" gorm:"embedded;embeddedPrefix:ticket_"`
	AllowedAppLanguages          Language        `json:"allowedAppLanguages" gorm:"embedded;embeddedPrefix:app_"`
	LanguageCode                 string          `json:"languageCode"`
	PrintTicket                  bool            `json:"printTicket"`
	OptionalPrintTicket          bool            `json:"optionalPrintTicket"`
	PrintTicketPrice             bool            `json:"printTicketPrice"`
	OptionalPrintTicketPrice     bool            `json:"optionalPrintTicketPrice"`
	PrintSummary                 bool            `json:"printSummary"`
	OptionalPrintSummary         bool            `json:"optionalPrintSummary"`
	HighlightPrintedReservations bool            `json:"highlightPrintedReservations"`
	SingleDocPrint               bool            `json:"singleDocPrint"`
	AllowedTicketGroupTypes      []string        `json:"allowedTicketGroupTypes" gorm:"type:JSON"`
	PrintCashCount               bool            `json:"printCashCount"`
	ShowChangeCalculator         int             `json:"showChangeCalculator"`
	AskCashForChange             bool            `json:"askCashForChange"`
	Calculator                   bool            `json:"calculator"`
	ShowDeferredPaymentMethod    bool            `json:"showDeferredPaymentMethod"`
	ValidationMethod             int             `json:"validationMethod"`
	VerifierId                   int             `json:"verifierId"`
	AllowConfigPrinter           bool            `json:"allowConfigPrinter"`
	AllowLoginHistory            bool            `json:"allowLoginHistory"`
	AllowChangePassword          bool            `json:"allowChangePassword"`
	AllowSellAndEditPastEvents   bool            `json:"allowSellAndEditPastEvents"`
	PrintMultiple                bool            `json:"printMultiple"`
	CloseDetailsPrintTickets     bool            `json:"closeDetailsPrintTickets"`
	CloseDetailsPrintReceipt     bool            `json:"closeDetailsPrintReceipt"`
	AvailabilityThreshold        int             `json:"availabilityThreshold"`
	SalesGroupSet                []SalesGroup    `json:"salesGroupSet" gorm:"many2many:boxoffice_salesgroups;"`
	CreatedAt                    time.Time       `json:"-"`
	UpdatedAt                    time.Time       `json:"-"`
	DeletedAt                    gorm.DeletedAt  `json:"-" gorm:"index"`
}
