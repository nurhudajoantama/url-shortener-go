package helper

type Response struct {
	Success bool
	Message string
	Errors  interface{}
	Data    interface{}
}

func CreateSuccessResponse(message string, data interface{}) Response {
	return Response{
		Success: true,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
}

func CreateErrorResponse(message string, errors interface{}) Response {
	return Response{
		Success: false,
		Message: message,
		Errors:  errors,
		Data:    nil,
	}
}
