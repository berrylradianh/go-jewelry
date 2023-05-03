package product_categories

type ProductCategoryResponse struct {
	ID   int `json:"id,omitempty"`
	Name string
}

func (ProductCategoryResponse) TableName() string {
	return "product_categories"
}
