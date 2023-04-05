package models

type Node struct {
	Key       string  `json:"key"`
	Value     string  `json:"value"`
	Next      *string `json:"next"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
