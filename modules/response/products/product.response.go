package products

type ProductResponse struct {
	ID                  int     `json:"id" form:"id"`
	Name                string  `json:"name" form:"name"`
	Price               float64 `json:"price" form:"price"`
	Stock               int     `json:"stock" form:"stock"`
	Image_url           string  `json:"image_url" form:"image_url"`
	Product_category_id int     `json:"-"`
	Product_material_id int     `json:"-"`
}

func (ProductResponse) TableName() string {
	return "products"
}
