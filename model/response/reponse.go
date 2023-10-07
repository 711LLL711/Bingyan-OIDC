package response

type response struct {
	Success bool        `json:"success"`
	Hint    string      `json:"hint"`
	Data    interface{} `json:"data"`
}

type UserResponse struct {
	UserID     int    `json:"userID"`
	UserName   string `json:"userName" `
	UserAvatar string `json:"userAvatar"`
}

// 定义常用错误response
var (
	UnauthorizedError   = MakeFailedResponse("Unauthorized")
	UnVerifiedError     = MakeFailedResponse("Unverified")
	InvalidInfoError    = MakeFailedResponse("Invalid information")
	DatabaseError       = MakeFailedResponse("Database error")
	TimeoutError        = MakeFailedResponse("Timeout")
	InternalServerError = MakeFailedResponse("Internal server error")
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
