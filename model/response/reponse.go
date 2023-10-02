package response

type response struct {
	Success bool        `json:"success"`
	Hint    string      `json:"hint"`
	Data    interface{} `json:"data"`
}

type UserResponse struct {
	// 比较通用的结构体，用户信息三要素
	// 场景：登陆/注册/更新/查询 用户信息后返回
	UserID     uint   `json:"userID"`
	UserName   string `json:"userName" `
	UserAvatar string `json:"userAvatar"`
}

// 定义常用错误
var (
	UnauthorizedError = MakeFailedResponse("Unauthorized")
	InvalidInfoError  = MakeFailedResponse("Invalid information")
	DatabaseError     = MakeFailedResponse("Database error")
	TimeoutError      = MakeFailedResponse("Timeout")
)

func MakeResponse(success bool, hint string, data interface{}) response {
	return response{
		Success: success,
		Hint:    hint,
		Data:    data,
	}
}

func MakeSucceedResponse(data interface{}) response {
	return MakeResponse(true, "", data)
}

func MakeFailedResponse(hint string) response {
	return MakeResponse(false, hint, nil)
}
