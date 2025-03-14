package exception

import "github.com/lambertstu/shortlink-core-rpc/pkg/errorcode"

type AbstractException struct {
	ErrorCode    string
	ErrorMessage string
	Err          error
}

func NewAbstractException(message string, code errorcode.ErrorCode, err error) *AbstractException {
	if message == "" {
		message = code.Message()
	}
	return &AbstractException{
		ErrorCode:    code.Code(),
		ErrorMessage: message,
		Err:          err,
	}
}
