package orders

import "time"

type Order struct {
	ID        int       `json:"id"`
	Total     int       `json:"total"`
	UserID    int       `json:"userID"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
