package products

type ProductResponse struct {
	Name                string
	Price               string
	Stock               int
	Image_url           string
	Product_category_id int `json:"product_category_id,omitempty"`
}

func (ProductResponse) TableName() string {
	return "products"
}
