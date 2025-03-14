package exception

import "github.com/lambertstu/shortlink-core-rpc/pkg/errorcode"

type ClientException struct {
	*AbstractException
}

// NewClientExceptionWithErr 创建一个没有消息的 RemoteException
func NewClientExceptionWithErr(errorCode errorcode.ErrorCode) *ClientException {
	exception := NewAbstractException("", errorCode, nil)
	return &ClientException{exception}
}

// NewClientExceptionWithMsg 创建一个带有消息的 RemoteException
func NewClientExceptionWithMsg(message string) *ClientException {
	exception := NewAbstractException(message, errorcode.ClientError, nil)
	return &ClientException{exception}
}

// NewClientException 创建一个带有消息和错误码的 RemoteException
func NewClientException(message string, code errorcode.ErrorCode, err error) *ClientException {
	exception := NewAbstractException(message, code, err)
	return &ClientException{exception}
}
