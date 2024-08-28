package lv_dto

type RespPage struct {
	// 代码
	Code  int   `json:"code" example:"200"`
	Total int64 `json:"total"`
	// 数据集
	Rows any `json:"rows"`
	// 消息
	Msg string `json:"msg"`
}

func FailPage(msg string) RespPage {
	return RespPage{Code: 500, Msg: msg}
}

func SuccessPage[T any](list any, total int64) RespPage {
	return RespPage{Code: 200, Rows: list, Total: total, Msg: "success"}
}
