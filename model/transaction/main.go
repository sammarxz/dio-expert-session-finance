package transaction

import "time"

type Transaction struct {
	Title      string    `json:"title"`
	Amount     float32   `json:"amount"`
	Type       int       `json:"type"`
	Created_at time.Time `json:"created_at"`
}

type Transactions []Transaction
