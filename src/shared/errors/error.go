package error

type Error struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Error      string `json:"error"`
}

func New(message string, statusCode int, err error) *Error {

	var messageErrror string = ""
	if err != nil {
		messageErrror = err.Error()
	}

	if message == "" || statusCode == 0 {
		return &Error{
			Message:    "Internal Server Error",
			StatusCode: 500,
			Error:      messageErrror,
		}
	}

	return &Error{
		Message:    message,
		StatusCode: statusCode,
		Error:      messageErrror,
	}
}
