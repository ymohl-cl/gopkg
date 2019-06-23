package jsonvalidator

import (
	"gopkg.in/go-playground/validator.v9"
)

// JSONValidator is an implementation of validation data stucture
// will be attached on http driver
type JSONValidator interface {
	Validate(i interface{}) error
}

type jsonValidator struct {
	driver *validator.Validate
}

// New return an implement of JSONValidator
func New() JSONValidator {
	return &jsonValidator{
		driver: validator.New(),
	}
}

// Validate check the validity structure
func (j jsonValidator) Validate(i interface{}) error {
	if err := j.driver.Struct(i); err != nil {
		return err
	}
	return nil
}
