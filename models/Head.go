package models

type Head struct {
	Key       string  `json:"key"`
	Next      *string `json:"next"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
