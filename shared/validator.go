package shared

import (
	"errors"
	"reflect"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTrans "github.com/go-playground/validator/v10/translations/en"
)

func Validator(data interface{}, rule string) []error {
	var err error
	validate := validator.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = enTrans.RegisterDefaultTranslations(validate, trans)

	if reflect.TypeOf(data).Kind() == reflect.Struct {
		err = validate.Struct(data)
	} else {
		err = validate.Var(data, rule)
	}

	if err == nil {
		return nil
	}

	var errs []error
	for _, e := range err.(validator.ValidationErrors) {
		errText := errors.New(e.Field() + ": " + e.Translate(trans))
		errs = append(errs, errText)
	}

	return errs
}
