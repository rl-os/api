package pkg

// ErrorResponse default struct for errors
type ErrorResponse struct {
	// contain short, unique error
	Error string `json:"error"`
	// full description about error
	ErrorDescription string `json:"error_description"`
	// addition information
	Message string `json:"message"`
}

// HTTPError contains Go-style error with information like description, message and etc
type HTTPError struct {
	Body ErrorResponse `json:"-"`
}

func (he *HTTPError) Error() string {
	return he.Body.ErrorDescription
}

// NewHTTPError return new Go-style error with loaded information from config file
func NewHTTPError(code int, err string, message string) error {
	return &HTTPError{
		Body: ErrorResponse{
			Error: err,
			ErrorDescription: message,
			Message: message,
		},
	}
}