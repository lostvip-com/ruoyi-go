package dtos

import "encoding/json"

type InvokeLog struct {
	Req      interface{} `json:"req"`
	Reply    interface{} `json:"reply"`
	Success  bool        `json:"success"`
	Method   string      `json:"method"`
	Duration string      `json:"duration"`
	Error    string      `json:"error"`
}

func (il InvokeLog) ToString() string {
	data, _ := json.Marshal(il)
	return string(data)
}

type HandleLog struct {
	Req      interface{} `json:"req"`
	Reply    interface{} `json:"reply"`
	Success  bool        `json:"success"`
	Method   string      `json:"method"`
	Duration string      `json:"duration"`
	Error    string      `json:"error"`
}

func (hl HandleLog) ToString() string {
	data, _ := json.Marshal(hl)
	return string(data)
}
