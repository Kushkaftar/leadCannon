package models

// Config ...
type Config struct {
	Main       Main       `json:"main"`
	Redirect   Redirect   `json:"redirect"`
	Publisher  Publisher  `json:"publisher"`
	Advertiser Advertiser `json:"advertiser"`
}

// Main ...
type Main struct {
	Run        bool   `json:"run"`
	URL        string `json:"url" validate:"required,url"`
	Sum        int    `json:"sum" validate:"required"`
	Label      bool   `json:"label" validate:"required"`
	Subaccaunt bool   `json:"subaccaunt" validate:"required"`
	Advertiser string `json:"advertiser" validate:"required"` // false, true, trash, redject, approved
}

// Publisher ...
type Publisher struct {
	API_Key string `json:"api_key" validate:"required"`
	Hash    string `json:"hash" validate:"required,uuid"`
}

// Advertiser ...
type Advertiser struct {
	API_Key string `json:"api_key" validate:"required"`
}

// Redirect ...
type Redirect struct {
	Run             bool   `json:"run"`
	Redirect_Domain string `json:"redirect_domain" validate:"required"`
	UUID            string `json:"uuid" validate:"required,uuid"`
	Sum             int    `json:"sum" validate:"required"`
}
