package utils

import (
	"github.com/astaxie/beego/validation"

	"ginIMApi/packages/logging"
)

func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}

	return
}
