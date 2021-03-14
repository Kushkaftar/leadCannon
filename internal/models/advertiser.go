package models

type Trash struct {
	OrderID     string `json:"order_id"`
	TrashReason int    `json:"trash_reason"`
	IsFraud     bool   `json:"is_fraud"`
}

type Redject struct {
	OrderID       string `json:"order_id"`
	RedjectReason int    `json:"reject_reason"`
}

type Approve struct {
	OrderID string `json:"order_id"`
}
