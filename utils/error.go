package utils

import "github.com/astaxie/beego/validation"

// ErrorsPack join all errors
func ErrorsPack(errors []*validation.Error) string {
	errString := ""
	for _, error := range errors {
		errString += error.Error()
	}
	return errString
}
