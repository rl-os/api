package oauth

// ErrorResponse default struct for errors
type ErrorResponse struct {
	// contain short, unique error
	Error string `json:"error"`
	// full description about error
	ErrorDescription string `json:"error_description"`
	// addition information
	Message string `json:"message"`
}
