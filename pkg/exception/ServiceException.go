package exception

import "github.com/lambertstu/shortlink-core-rpc/pkg/errorcode"

type ServiceException struct {
	*AbstractException
}

// NewServiceExceptionWithErr 创建一个没有消息的 ServiceException
func NewServiceExceptionWithErr(errorCode errorcode.ErrorCode) *ServiceException {
	exception := NewAbstractException("", errorCode, nil)
	return &ServiceException{exception}
}

// NewServiceExceptionWithMsg 创建一个带有消息的 ServiceException
func NewServiceExceptionWithMsg(message string) *ServiceException {
	exception := NewAbstractException(message, errorcode.ServiceError, nil)
	return &ServiceException{exception}
}

// NewServiceException 创建一个带有消息和错误码的 ServiceException
func NewServiceException(message string, code errorcode.ErrorCode, err error) *ServiceException {
	exception := NewAbstractException(message, code, err)
	return &ServiceException{exception}
}
