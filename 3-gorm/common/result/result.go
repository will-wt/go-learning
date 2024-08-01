package result

// Any 泛型接口
type Any interface{}

// Result API统一返回结果，使用 omitempty 标签来排除 null 值的字段
type Result struct {
	Code    uint16 `json:"code"`
	Message string `json:"message,omitempty"`
	Data    Any    `json:"data,omitempty"`
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
