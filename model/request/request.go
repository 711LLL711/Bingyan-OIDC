package request

// UserRegisterRequest 注册时的请求
type UserRegisterRequest struct {
	Username string `json:"username" validate:"required" form:"username" `
	Email    string `json:"email" validate:"required" form:"email"`
	Password string `json:"password" validate:"required" form:"password" `
}

// UserSignInRequest 登陆时的请求
type UserLogInRequest struct {
	Email    string `json:"email" validate:"required" form:"email"`
	Password string `json:"password" validate:"required" form:"password" `
}

// UserUpdateRequest 更新用户信息时的请求
type UserUpdateRequest struct {
	UserID   uint   `json:"-"` // 在处理请求时，从 token 中获取；controller 负责赋值
	UserName string `validate:"omitempty" json:"userName,omitempty"`
	Bio      string `validate:"omitempty,max=100" json:"bio,omitempty"`
}
