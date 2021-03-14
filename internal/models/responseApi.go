package models

type Response struct {
	Success bool `json:"success"`
}

type ErrorResponse struct {
	Errors []Error `json:"errors"`
}

type Error struct {
	Text  string `json:"text"`
	Field string `json:"field"`
}

type DataResponse struct {
	Data Data `json:"data"`
}

type Data struct {
	LeadUUID string `json:"lead_uuid"`
}
