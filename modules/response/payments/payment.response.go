package payments

type PaymentResponse struct {
	ID   string `json:"-"`
	Name string `json:"name,omitempty" form:"name"`
}

func (PaymentResponse) TableName() string {
	return "payments"
}
