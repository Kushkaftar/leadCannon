package models

// RedirectLead ...
type RedirectLead struct {
	RID         string `json:"r_id"`
	CountryCode string `json:"country_code"`
	Phone       string `json:"phone"`
	Name        string `json:"name"`
	Comment     string `json:"comment"`
	OrderAmount string `json:"order_amount"`
}
