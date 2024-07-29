package result

// Any 泛型接口
type Any interface {
}

// Result API统一返回结果
type Result struct {
	Code    uint16 `json:"code"`
	Message string `json:"message"`
	Data    Any    `json:"data"`
}

func Ok() Result {
	return Result{
		Code:    200,
		Message: "success",
	}
}

func OkWithData(data Any) Result {
	return Result{
		Code:    200,
		Message: "success",
		Data:    data,
	}
}

func Error(code uint16, message string) Result {
	return Result{
		Code:    code,
		Message: message,
	}
}

func Of(code uint16, message string, data Any) Result {
	return Result{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
