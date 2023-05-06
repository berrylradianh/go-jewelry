package users

type UserDetailResponse struct {
	Name    string `json:"name,omitempty" form:"name"`
	Address string `json:"address,omitempty" form:"address"`
	Gender  string `json:"gender,omitempty" form:"gender"`
	Phone   string `json:"phone,omitempty" form:"phone"`
	User_id int    `json:"-"`
}

func (UserDetailResponse) TableName() string {
	return "user_details"
}
