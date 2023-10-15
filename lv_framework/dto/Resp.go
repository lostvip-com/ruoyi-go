package dto

type Resp struct {
	// 代码
	Code int `json:"code" example:"200"`
	// 数据集
	Data interface{} `json:"data"`
	// 消息
	Msg string `json:"msg"`
}

func (res *Resp) ReturnOK() *Resp {
	res.Code = 200
	res.Msg = "Success!"
	return res
}
func (res *Resp) RetData(data *interface{}) *Resp {
	res.Code = 200
	res.Data = data
	return res
}
func (res *Resp) ReturnError(code int) *Resp {
	res.Code = code
	return res
}

func (res *Resp) Fail(msg string) *Resp {
	res.Code = 1
	res.Msg = msg
	return res
}

func (res *Resp) Ok(data *interface{}) *Resp {
	res.Code = 200
	res.Data = data
	res.Msg = "success!"
	return res
}
