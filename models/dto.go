package models

type ResultDto struct {
	Msg  string
	Data interface{}
	Code int
}

func SuccessResult(data interface{}) *ResultDto {
	return &ResultDto{Msg: "success", Data: data, Code: 200}
}

func FailResult(data interface{}) *ResultDto {
	return &ResultDto{Msg: "fail", Data: data, Code: 999}
}
