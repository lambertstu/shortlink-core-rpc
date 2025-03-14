package errorcode

type BaseErrorCode struct {
	code    string
	message string
}

func NewBaseErrorCode(code, message string) *BaseErrorCode {
	return &BaseErrorCode{
		code:    code,
		message: message,
	}
}

func (b *BaseErrorCode) Code() string {
	return b.code
}

func (b *BaseErrorCode) Message() string {
	return b.message
}

var (
	// ========== 一级宏观错误码 客户端错误 ==========
	ClientError = NewBaseErrorCode("A000001", "用户端错误")

	// ========== 二级宏观错误码 用户注册错误 ==========
	UserRegisterError             = NewBaseErrorCode("A000100", "用户注册错误")
	UserNameVerifyError           = NewBaseErrorCode("A000110", "用户名校验失败")
	UserNameExistError            = NewBaseErrorCode("A000111", "用户名已存在")
	UserNameSensitiveError        = NewBaseErrorCode("A000112", "用户名包含敏感词")
	UserNameSpecialCharacterError = NewBaseErrorCode("A000113", "用户名包含特殊字符")
	UserNotExist                  = NewBaseErrorCode("A000114", "用户不存在")
	PasswordVerifyError           = NewBaseErrorCode("A000120", "密码校验失败")
	PasswordShortError            = NewBaseErrorCode("A000121", "密码长度不够")
	PhoneVerifyError              = NewBaseErrorCode("A000151", "手机格式校验失败")

	// ========== 二级宏观错误码 系统请求缺少幂等Token ==========
	IdempotentTokenNullError   = NewBaseErrorCode("A000200", "幂等Token为空")
	IdempotentTokenDeleteError = NewBaseErrorCode("A000201", "幂等Token已被使用或失效")

	// ========== 一级宏观错误码 系统执行出错 ==========
	ServiceError        = NewBaseErrorCode("B000001", "系统执行出错")
	ServiceTimeoutError = NewBaseErrorCode("B000100", "系统执行超时")

	// ========== 一级宏观错误码 调用第三方服务出错 ==========
	RemoteError = NewBaseErrorCode("C000001", "调用第三方服务出错")
)
