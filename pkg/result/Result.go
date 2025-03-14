package result

const SUCCESS_CODE = "0"

// Result represents a global response object with generic data type.
type Result[T any] struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
	Success bool   `json:"success"`
}

func NewResult[T any]() *Result[T] {
	return &Result[T]{}
}

func (r *Result[T]) IsSuccess() bool {
	return r.Code == SUCCESS_CODE
}

func (r *Result[T]) SetCode(code string) *Result[T] {
	r.Code = code
	return r
}

func (r *Result[T]) SetMessage(message string) *Result[T] {
	r.Message = message
	return r
}

func (r *Result[T]) SetData(data T) *Result[T] {
	r.Data = data
	return r
}
