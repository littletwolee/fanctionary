package models

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func NewResult(err error, data interface{}) Result {
	var (
		code int
		msg  string
	)
	if err != nil {
		code = 1
		msg = err.Error()
	}
	return Result{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
