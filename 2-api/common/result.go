package common

// Any 泛型接口
type Any interface {
}

// Result API统一返回结果
type Result struct {
	Code    uint16 `json:"code"`
	Message string `json:"message"`
	Data    Any    `json:"data"`
}

func OfSuccessResult(data Any) Result {
	return Result{
		Code:    200,
		Message: "success",
		Data:    data,
	}
}

func OfResult(code uint16, message string, data Any) Result {
	return Result{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
