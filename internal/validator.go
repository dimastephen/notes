package validator

import (
	"strings"
	"unicode/utf8"
)

type Validator struct {
	FieldsErrors map[string]string
}

func (v *Validator) Valid() bool {
	if len(v.FieldsErrors) == 0 {
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

func PermittedInt(value int, permittedValues ...int) bool {
	for i := range permittedValues {
		if permittedValues[i] == value {
			return true
		}
	}
	return false
}
