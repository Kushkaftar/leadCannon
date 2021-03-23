package models

// Lead ...
type Lead struct {
	CountryCode string `json:"country_code"`
	Phone       string `json:"phone"`
	Name        string `json:"name"`
	IP          string `json:"ip"`
	FlowUUID    string `json:"flow_uuid"`
	SubAccount  string `json:"subaccount"`
	UserAgent   string `json:"user_agent"`
	UtmCampaign string `json:"utm_campaign"`
	UtmContent  string `json:"utm_content"`
	UtmTerm     string `json:"utm_term"`
	UtmSource   string `json:"utm_source"`
	UtmMedium   string `json:"utm_medium"`
	SubID1      string `json:"sub_id1"`
	SubID2      string `json:"sub_id2"`
	SubID3      string `json:"sub_id3"`
	Referer     string `json:"referer"`
	Comment     string `json:"comment"`
}
