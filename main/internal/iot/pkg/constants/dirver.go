package constants

const (
	DefaultDriverBaseAddress = "localhost:-1"
	HostAddress              = "127.0.0.1"

	ConfigSuffix = ".toml"

	ConfigKeyDriver = "driver"

	DriverBaseDir            = "driver-data" // 驱动相关目录dir
	DriverLibraryDir         = DriverBaseDir + "/driver-library"
	DriverBinDir             = DriverBaseDir + "/bin"
	DriverRunConfigDir       = DriverBaseDir + "/run-config"
	DriverMntDir             = DriverBaseDir + "/mnt"
	DriverDefaultLogPath     = "logs/driver.log"
	DockerHummingbirdRootDir = "/var/bin/hummingbird"
)

const (
	DriverLibTypeDefault = iota + 1
	DriverLibTypeAppService
)

const (
	DeviceLibraryUploadTypeConfig = 2

	// 驱动配置定义，key=>value 的type
	DriverConfigTypeInt    = "int"
	DriverConfigTypeFloat  = "float"
	DriverConfigTypeString = "string"
	DriverConfigTypeBool   = "bool"
	DriverConfigTypeSelect = "select"
	DriverConfigTypeObject = "object"
	DriverConfigTypeArray  = "array"
)
