package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(New)

// New validator instance
func New() (*validator.Validate, error) {
	return validator.New(), nil
}
