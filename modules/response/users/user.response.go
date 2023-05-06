package users

type UserResponse struct {
	ID      int    `json:"-"`
	Email   string `json:"email,omitempty" form:"email"`
	Role_id int    `json:"-"`
}

func (UserResponse) TableName() string {
	return "users"
}
