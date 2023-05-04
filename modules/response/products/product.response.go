package products

type ProductResponse struct {
	ID                     int     `json:"id,omitempty" form:"id"`
	Name                   string  `json:"name,omitempty" form:"name"`
	Price                  float64 `json:"price,omitempty" form:"price"`
	Stock                  int     `json:"stock,omitempty" form:"stock"`
	Image_url              string  `json:"image_url,omitempty" form:"image_url"`
	Product_category_id    int     `json:"-"`
	Product_material_id    int     `json:"-"`
	Product_description_id int     `json:"-"`
}

func (ProductResponse) TableName() string {
	return "products"
}
