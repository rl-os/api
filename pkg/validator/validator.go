package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(New)

type Inst struct {
	*validator.Validate
}

// New validator instance
func New() (*Inst, error) {
	return &Inst{validator.New()}, nil
}
