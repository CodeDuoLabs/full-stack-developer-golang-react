package response

type Response struct {
	Status int    `json:"status"`
	Data   any    `json:"data"`
	Error  string `json:"error"`
}

func NewSuccessResponse(status int, data interface{}) Response {
	return Response{
		Status: status,
		Data:   data,
		Error:  "",
	}
}

func NewErrorResponse(status int, errorMsg string) Response {
	return Response{
		Status: status,
		Data:   nil,
		Error:  errorMsg,
	}
}
