package roles

type RoleResponse struct {
	ID   string `json:"-"`
	Name string `json:"name" form:"name"`
}

func (RoleResponse) TableName() string {
	return "roles"
}
