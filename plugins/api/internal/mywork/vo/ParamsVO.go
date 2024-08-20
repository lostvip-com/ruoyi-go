package vo

type ParamsVO struct {
	ParamNo   int    `form:"paramNo"`
	ParamName string `form:"paramName"binding:"required" `
}
