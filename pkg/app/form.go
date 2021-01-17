package app

import (
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"strings"
)

type (
	ValidError struct {
		Key     string
		Message string
	}

	ValidErrors []*ValidError

	CustomValidator struct {
		validator *val.Validate
	}
)

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func NewValidator() *CustomValidator {
	return &CustomValidator{validator: val.New()}
}

func BindAndValid(e echo.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	if err := e.Bind(v); err != nil {
		v := e.Get("trans")
		trans, _ := v.(ut.Translator)
		verrs, ok := err.(val.ValidationErrors)
		if !ok {
			return false, nil
		}

		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}
		return false, errs
	}

	if err := e.Validate(v); err != nil {
		return false, ValidErrors{&ValidError{Key: "", Message: err.Error()}}
	}

	return true, nil
}
