package validations

import (
	"github.com/gin-gonic/gin/binding"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Validator struct{}

func (v Validator) Validate(i interface{}) error {
	if c, ok := i.(validation.Validatable); ok {
		return c.Validate()
	}
	return nil
}

type BindingValidator struct {
	validator *Validator
}

func (bv *BindingValidator) ValidateStruct(obj any) error {
	return bv.validator.Validate(obj)
}

func (bv *BindingValidator) Engine() any {
	return bv.validator
}

func NewBindingValidator() binding.StructValidator {
	return &BindingValidator{validator: &Validator{}}
}
