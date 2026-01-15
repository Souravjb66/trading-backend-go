package models

type Order struct {
	ID                int64   `json:"id"`
	UserID            int64   `json:"user_id"`
	Asset             string  `json:"asset"`
	Type              string  `json:"type"`
	Price             float64 `json:"price"`
	Quantity          float64 `json:"quantity"`
	RemainingQuantity float64 `json:"remaining_quantity"`
	Status            string  `json:"status"`
	CreatedAt         string  `json:"created_at"`
}