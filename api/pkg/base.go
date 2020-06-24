package pkg

// ErrorResponse default struct for errors
type ErrorResponse struct {
	// contain short, unique errors
	ErrorID string `json:"error"`
	// full description about errors
	ErrorDescription string `json:"error_description"`
	// addition information
	Message string `json:"message"`

	// response status
	Status int `json:"-"`
}

func (he *ErrorResponse) Error() string {
	return he.ErrorDescription
}

// NewHTTPError return new Go-style errors with loaded information from config file
func NewHTTPError(code int, err string, message string) error {
	return &ErrorResponse{
		ErrorID:          err,
		ErrorDescription: message,
		Message:          message,
		Status:           code,
	}
}
