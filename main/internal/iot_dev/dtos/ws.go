package dtos

type WsCode uint32

const (
	WsCodeDeviceLibraryUpgrade   WsCode = 10001 // 驱动下载/升级
	WsCodeDeviceServiceRunStatus WsCode = 10002 // 驱动重启
	WsCodeDeviceLibraryDelete    WsCode = 10003 // 驱动删除
	WsCodeDeviceServiceLog       WsCode = 10004 // 驱动日志

	WsCodeCheckLang WsCode = 30001 // 切换语言

)
