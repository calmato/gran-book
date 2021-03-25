package validation

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/calmato/gran-book/api/server/book/internal/domain/exception"
	"github.com/go-playground/validator/v10"
)

// RequestValidator - リクエストバリデーションインターフェース
type RequestValidator interface {
	Run(i interface{}) []*exception.ValidationError
}

type requestValidator struct {
	validate validator.Validate
}

// NewRequestValidator - Validatorの生成
func NewRequestValidator() RequestValidator {
	validate := validator.New()

	if err := validate.RegisterValidation("hiragana", hiraganaCheck); err != nil {
		return nil
	}

	if err := validate.RegisterValidation("password", passwordCheck); err != nil {
		return nil
	}

	return &requestValidator{
		validate: *validate,
	}
}

const (
	hiraganaString = "^[ぁ-ゔー]*$"
	passwordString = "^[a-zA-Z0-9_!@#$_%^&*.?()-=+]*$"
)

var (
	hiraganaRegex = regexp.MustCompile(hiraganaString)
	passwordRegex = regexp.MustCompile(passwordString)
)

// Run - バリデーションの実行
func (rv *requestValidator) Run(i interface{}) []*exception.ValidationError {
	err := rv.validate.Struct(i)
	if err == nil {
		return []*exception.ValidationError{}
	}

	errors := err.(validator.ValidationErrors)
	validationErrors := make([]*exception.ValidationError, len(errors))

	rt := reflect.ValueOf(i).Elem().Type()

	for i, v := range errors {
		errorField, _ := rt.FieldByName(v.Field())
		errorFieldName := errorField.Tag.Get("json")
		errorMessage := ""

		// TODO: ネストしてるとこ、うまくjsonタグの値が取れないため
		if errorFieldName == "" {
			errorFieldName = v.Field()
		}

		switch v.Tag() {
		case exception.EqFieldTag:
			eqField, _ := rt.FieldByName(v.Param())
			errorMessage = validationMessage(v.Tag(), eqField.Tag.Get("json"))
		default:
			errorMessage = validationMessage(v.Tag(), v.Param())
		}

		validationErrors[i] = &exception.ValidationError{
			Field:   errorFieldName,
			Message: errorMessage,
		}
	}

	return validationErrors
}

func hiraganaCheck(fl validator.FieldLevel) bool {
	return hiraganaRegex.MatchString(fl.Field().String())
}

func passwordCheck(fl validator.FieldLevel) bool {
	return passwordRegex.MatchString(fl.Field().String())
}

func validationMessage(tag string, options ...string) string {
	switch tag {
	case exception.RequiredTag:
		return exception.RequiredMessage
	case exception.EqFieldTag:
		return fmt.Sprintf(exception.EqFieldMessage, options[0])
	case exception.MinTag:
		return fmt.Sprintf(exception.MinMessage, options[0])
	case exception.MaxTag:
		return fmt.Sprintf(exception.MaxMessage, options[0])
	case exception.Base64Tag:
		return exception.Base64Message
	case exception.EmailTag:
		return exception.EmailMessage
	case exception.HiraganaTag:
		return exception.HiraganaMessage
	case exception.PasswordTag:
		return exception.PasswordMessage
	case exception.UniqueTag:
		return exception.UniqueMessage
	case exception.LessThanTag:
		return fmt.Sprintf(exception.LessThanMessage, options[0])
	case exception.GreaterThanTag:
		return fmt.Sprintf(exception.GreaterThanMessage, options[0])
	case exception.LessThanEqualTag:
		return fmt.Sprintf(exception.LessThanEuqalMessage, options[0])
	case exception.GreaterThanEqualTag:
		return fmt.Sprintf(exception.GreaterThanEqualMessage, options[0])
	case exception.OneOfTag:
		str := strings.ReplaceAll(options[0], " ", ",")
		return fmt.Sprintf(exception.OneOfMessage, str)
	default:
		return "Unknown"
	}
}
