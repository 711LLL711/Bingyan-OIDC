package request

// UserRegisterRequest 注册时的请求
type UserRegisterRequest struct {
	Username string `json:"username" validate:"required" `
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required" `
}

// UserSignInRequest 登陆时的请求
type UserLogInRequest struct {
	Email    string `json:"email" validate:"required" `
	Password string `json:"password" validate:"required" `
}

// UserUpdateRequest 更新用户信息时的请求
type UserUpdateRequest struct {
	UserID   uint   `json:"-"` // 在处理请求时，从 token 中获取；controller 负责赋值
	UserName string `validate:"omitempty,min=3,max=20" json:"userName,omitempty"`
	Password string `validate:"omitempty,min=8,max=20" json:"newPassword,omitempty"`
	Avatar   string `validate:"omitempty,url" json:"avatar,omitempty"`
	Bio      string `validate:"omitempty,max=100" json:"bio,omitempty"`
}
