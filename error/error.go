package error

type Error struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func New(message string, statusCode int) *Error {
	if message == "" || statusCode == 0 {
		return &Error{
			Message:    "Internal Server Error",
			StatusCode: 500,
		}

	}

	return &Error{
		Message:    message,
		StatusCode: statusCode,
	}
}
