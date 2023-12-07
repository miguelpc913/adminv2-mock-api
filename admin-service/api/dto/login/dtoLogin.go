package dtoLogin

type LoginReq struct {
	UserName  string `json:"username"`
	Password  string `json:"password"`
	Module    string `json:"module"`
	GrantType string `json:"grant_type"`
}
