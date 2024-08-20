package dtos

import "encoding/json"

type RpcData struct {
	Code    WsCode
	ReqId   string
	ErrCode uint32
	Data    interface{}
}

func (rd *RpcData) String() string {
	str, _ := json.Marshal(rd)
	return string(str)
}
