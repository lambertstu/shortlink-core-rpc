package errorcode

// ErrorCode 平台错误码｜定义错误码抽象接口，由各错误码类实现接口方法
type ErrorCode interface {
	// Code 错误码
	Code() string
	// Message 错误信息
	Message() string
}
