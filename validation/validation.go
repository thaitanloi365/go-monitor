package validation

import (
	"fmt"
	"net/http"

	english "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/en"
	"github.com/thaitanloi365/go-monitor/errs"
)

// Validator instance
type Validator struct {
	validator *validator.Validate
	trans     ut.Translator
}

// Validate validate
func (v *Validator) Validate(i interface{}) error {
	err := v.validator.Struct(i)
	if err == nil {
		return nil
	}

	if ers, ok := err.(validator.ValidationErrors); ok {
		var e = ValidateErrors{
			ers,
		}

		var err = e.TransformError().WithTranslate(v.trans)
		var transformError = getErrorForTag(err)
		if transformError != nil {
			return transformError
		}

		return errs.New(http.StatusUnprocessableEntity, err.Error())

	}

	return nil
}

// RegisterValidation register
func (v *Validator) RegisterValidation(tag string, fc validator.Func, msg string) error {
	err := v.validator.RegisterValidation(tag, fc)
	if err != nil {
		return fmt.Errorf("RegisterValidation error for %s", tag)
	}

	err = v.validator.RegisterTranslation(tag, v.trans, registrationFunc(tag, msg), translateFunc)
	if err != nil {
		return fmt.Errorf("RegisterTranslation error for %s", tag)
	}

	return nil
}

// RegisterValidation override echo's validator
func RegisterValidation() *Validator {
	v := validator.New()
	eng := english.New()
	uni := ut.New(eng, eng)
	trans, _ := uni.GetTranslator("en")

	customValidator := &Validator{validator: v, trans: trans}

	err := en.RegisterDefaultTranslations(customValidator.validator, trans)
	if err != nil {
		panic(err)
	}

	err = customValidator.validator.RegisterTranslation("required", customValidator.trans, registrationFunc("required", "{0} is required"), translateFunc)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	err = customValidator.RegisterValidation("isURL", isURL, fmt.Sprintf("{0} is invalid URL"))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	err = customValidator.RegisterValidation("isBirthday", isBirthday, fmt.Sprintf("{0} is invalid birthday"))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	err = customValidator.RegisterValidation("isTimezone", isTimezone, fmt.Sprintf("{0} is invalid timezone"))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	err = customValidator.RegisterValidation("isPhone", isPhone, fmt.Sprintf("{0} is invalid phone"))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return customValidator
}

func registrationFunc(tag string, translation string) validator.RegisterTranslationsFunc {
	return func(ut ut.Translator) error {
		return ut.Add(tag, translation, true)
	}

}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {
	t, err := ut.T(fe.Tag(), fe.Field())

	if err != nil {
		return fe.(error).Error()
	}
	return t
}

func getErrorForTag(err *ValidateError) error {
	fmt.Printf("Validation error FieldName = %s ValidateType = %s ExpectedResult = %s GotResult = %v\n", err.FieldName, err.ValidateType, err.ExpectedResult, err.GotResult)
	if err.ValidateType != "required" {
		return nil
	}

	return nil
}
