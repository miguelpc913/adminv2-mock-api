package appUserDto

type UserType string
type UserProfile string

const (
	UserTypeCallCenter UserType = "call_center"
	UserTypeBoxOffice  UserType = "box_office"

	UserProfileOperator UserProfile = "operator"
	UserProfileAdmin    UserProfile = "admin"
)

type AppUserCreateDTO struct {
	Status         bool        `json:"status"`
	Type           UserType    `json:"type"`
	Profile        UserProfile `json:"profile"`
	Name           string      `json:"name"`
	LastName       string      `json:"lastName"`
	UserName       string      `json:"username"` // maps to model field "UserName"
	Email          string      `json:"email"`
	Password       string      `json:"password"`
	PointOfSaleSet []int       `json:"pointOfSaleSet"` // IDs of POS sets
	ReportSet      []int       `json:"reportSet"`      // IDs of Report sets
}

type AppUserIdentityDTO struct {
	Status   bool        `json:"status"`
	Profile  UserProfile `json:"profile"`
	Name     string      `json:"name"`
	LastName string      `json:"lastName"`
	Email    string      `json:"email"`
}
