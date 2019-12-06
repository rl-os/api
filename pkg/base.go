package pkg

// ErrorResponse default struct for errors
type ErrorResponse struct {
	// contain short, unique error
	ErrorID string `json:"error"`
	// full description about error
	ErrorDescription string `json:"error_description"`
	// addition information
	Message string `json:"message"`
}

func (he *ErrorResponse) Error() string {
	return he.ErrorDescription
}

// NewHTTPError return new Go-style error with loaded information from config file
func NewHTTPError(code int, err string, message string) error {
	return &ErrorResponse{
		ErrorID: err,
		ErrorDescription: message,
		Message: message,
	}
}