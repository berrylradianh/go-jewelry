package products

type ProductMaterialResponse struct {
	ID   int    `json:"-"`
	Name string `json:"name" form:"name"`
}

func (ProductMaterialResponse) TableName() string {
	return "product_materials"
}
