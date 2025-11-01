package errorx

// 110 代表用户系统
var PramsError = New(1101111, "参数错误")

type BizErr struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func New(code int, msg string) *BizErr {
	return &BizErr{Code: code, Msg: msg}
}

func (e *BizErr) Error() string {
	return e.Msg
}

type ErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *BizErr) Data() *ErrorResponse {
	return &ErrorResponse{Code: e.Code, Msg: e.Msg}
}
