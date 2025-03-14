package exception

import "github.com/lambertstu/shortlink-core-rpc/pkg/errorcode"

type RemoteException struct {
	*AbstractException
}

// NewRemoteExceptionWithErr 创建一个没有消息的 RemoteException
func NewRemoteExceptionWithErr(errorCode errorcode.ErrorCode) *RemoteException {
	exception := NewAbstractException("", errorCode, nil)
	return &RemoteException{exception}
}

// NewRemoteExceptionWithMsg 创建一个带有消息的 RemoteException
func NewRemoteExceptionWithMsg(message string) *RemoteException {
	exception := NewAbstractException(message, errorcode.RemoteError, nil)
	return &RemoteException{exception}
}

// NewRemoteException 创建一个带有消息和错误码的 RemoteException
func NewRemoteException(message string, code errorcode.ErrorCode, err error) *RemoteException {
	exception := NewAbstractException(message, code, err)
	return &RemoteException{exception}
}
