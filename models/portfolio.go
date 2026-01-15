package models

type Portfolio struct {
	ID        int64   `json:"id"`
	UserID    int64   `json:"user_id"`
	Asset     string  `json:"asset"`
	Balance   float64 `json:"balance"`
	UpdatedAt string  `json:"updated_at"`
}