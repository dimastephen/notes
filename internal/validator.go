package validator

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

type Validator struct {
	FieldsErrors   map[string]string
	NonFieldErrors []string
}

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func MinChar(value string, n int) bool {
	return utf8.RuneCountInString(value) >= n
}

func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

func (v *Validator) Valid() bool {
	if len(v.FieldsErrors) == 0 && len(v.NonFieldErrors) == 0 {
		return true
	} else {
		return false
	}
}

func (v *Validator) AddFieldError(key string, value string) {
	if v.FieldsErrors == nil {
		v.FieldsErrors = make(map[string]string)
	}
	_, exists := v.FieldsErrors[key]
	if !exists {
		v.FieldsErrors[key] = value
	}
}

func (v *Validator) AddNonFieldError(str string) {
	v.NonFieldErrors = append(v.NonFieldErrors, str)
}

func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddFieldError(key, message)
	}
}

func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""

}

func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) < n
}

func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	for i := range permittedValues {
		if permittedValues[i] == value {
			return true
		}
	}
	return false
}
