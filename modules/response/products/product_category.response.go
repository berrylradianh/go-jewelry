package products

type ProductCategoryResponse struct {
	ID   int    `json:"-"`
	Name string `json:"name" form:"name"`
}

func (ProductCategoryResponse) TableName() string {
	return "product_categories"
}
