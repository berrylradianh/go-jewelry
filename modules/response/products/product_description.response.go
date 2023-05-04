package products

type ProductDescriptionResponse struct {
	ID          int    `json:"-"`
	Description string `json:"description" form:"description"`
}

func (ProductDescriptionResponse) TableName() string {
	return "product_descriptions"
}
