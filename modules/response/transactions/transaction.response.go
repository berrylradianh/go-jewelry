package transactions

import "time"

type TransactionResponse struct {
	ID              int       `json:"-"`
	Date            time.Time `json:"date" form:"date"`
	Status          string    `json:"status" form:"status"`
	Image_proof_url string    `json:"image_proof_url" form:"image_proof_url"`
	User_id         int       `json:"-" form:"-"`
	Payment_id      int       `json:"-" form:"-"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
