package orders

import "time"

type Order struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userID"`
	Status    status    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
