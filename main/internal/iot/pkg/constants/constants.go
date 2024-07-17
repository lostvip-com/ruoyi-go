package constants

var (
	//物模型能力类型
	Property = "property"
	Event    = "event"
	Action   = "action"
)

const (
	//gateway
	JwtParsedInfo          = "jwt-parsed-info"
	DefaultAgentReqTimeout = 10
)

// 默认高级配置ID
const DefaultAdvanceConfigID = 1

// 日志级别
type LogLevel int32

const (
	DebugLevel LogLevel = 0
	InfoLevel           = 1
	WarnLevel           = 2
	ErrorLevel          = 3
)

type DeviceServiceStatusType int32

const (
	// 驱动运行状态
	RunStatusStarted = iota + 1
	RunStatusStopped
	RunStatusStarting
	RunStatusStopping
)

// 驱动实例日志
const (
	StatusRead = iota + 1
	StatusStop
)

const (
	// 驱动库操作状态
	OperateStatusDefault    = "default"    // 默认（未安装）
	OperateStatusInstalling = "installing" // 安装中
	OperateStatusInstalled  = "installed"  // 已安装
	OperateStatusUninstall  = "uninstall"  // 未安装
)
