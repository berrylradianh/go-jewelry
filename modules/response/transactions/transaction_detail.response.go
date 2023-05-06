package transactions

type TransactionDetailResponse struct {
	Qty            int16   `json:"qty" form:"qty"`
	Price          float64 `json:"price" form:"price"`
	Transaction_id int     `json:"-"`
}

func (TransactionDetailResponse) TableName() string {
	return "transaction_details"
}
