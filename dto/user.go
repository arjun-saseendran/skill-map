package dto

type UserCreateInput struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
}

type UserUpdateInput struct {
	FullName stirng `json:"fullname"`
	Email    string `json:"email"`
}

func NewCreateUserInput() *UserCreateInput {
	return &UserCreateInput{}
}

func NewUserUpdateInput() *UserUpdateInput {
	return &UserUpdateInput{}
}
