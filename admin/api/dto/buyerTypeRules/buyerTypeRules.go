package dtoBuyerTypeRules

type VarsDTO struct {
	X []int `json:"x"`
	Y []int `json:"y,omitempty"`
	M *int  `json:"m,omitempty"`
	N *int  `json:"n,omitempty"`
}

type BuyerTypeRulePostDTO struct {
	Status                  bool    `json:"status"`
	Name                    string  `json:"name"`
	BuyerTypeRuleTemplateID uint    `json:"buyerTypeRuleTemplateId"`
	Vars                    VarsDTO `json:"vars"`
	ErrorMessage            string  `json:"errorMessage"`
	Priority                int     `json:"priority"`
	ProductSet              []int   `json:"productSet"`
}

type BuyerTypeRulesIdentity struct {
	Status bool   `json:"status"`
	Name   string `json:"name"`
}

type BuyerTypeRuleConfiguration struct {
	Vars                    VarsDTO `json:"vars"`
	ErrorMessage            string  `json:"errorMessage"`
	Priority                int     `json:"priority"`
	BuyerTypeRuleTemplateID uint    `json:"buyerTypeRuleTemplateId"`
}
