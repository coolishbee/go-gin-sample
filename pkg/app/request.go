package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/jameschun7/go-gin-sample/pkg/logging"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}

	return
}
