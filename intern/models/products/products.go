package products

import "time"

type Product struct {
	ID          int       `json:"ID"`
	Price       int       `json:"price"`
	Description string    `json:"description"`
	Created_at  time.Time `json:"created_at"`
}
