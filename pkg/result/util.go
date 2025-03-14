package result

import "github.com/lambertstu/shortlink-core-rpc/pkg/errorcode"

func Success() *Result[any] {
	return NewResult[any]().SetCode(SUCCESS_CODE)
}

func SuccessWithMsg[T any](data T) *Result[T] {
	return NewResult[T]().SetCode(SUCCESS_CODE).SetData(data)
}

func Failure() *Result[any] {
	serviceErr := errorcode.ServiceError
	return NewResult[any]().SetCode(serviceErr.Code()).SetData(serviceErr.Message())
}

func FailureWithMsg[T any](data T) *Result[any] {
	serviceErr := errorcode.ServiceError
	return NewResult[any]().SetCode(serviceErr.Code()).SetData(data)
}

func FailureWithErr[T any](data T) *Result[any] {
	serviceErr := errorcode.ServiceError
	return NewResult[any]().SetCode(serviceErr.Code()).SetData(data)
}
