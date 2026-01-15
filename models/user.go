package models


type User struct {
	ID       int64   `json:"id"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	// relations
	Portfolio  []Portfolio `json:"portfolio,omitempty"`
	OpenOrders []Order     `json:"openOrders,omitempty"`
}