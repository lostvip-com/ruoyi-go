package dtos

type (
	// LogConfig logger common
	LogConfig struct {
		FileName string
		LogLevel string
	}

	CloudInstanceLogConfig struct {
		FileName  string
		LogSwitch bool
		LogFilter []string
	}

	// RPCConfig internal grpc server common
	RPCConfig struct {
		Address  string
		UseTLS   bool
		CertFile string
		KeyFile  string
	}

	// ClientInfo provides the host and port of another service in tedge.
	ClientInfo struct {
		Address string
		// 是否启用tls
		UseTLS bool
		// ca cert
		CertFilePath string
		// mqtt clientId
		ClientId string
		// mqtt username
		Username string
		// mqtt password
		Password string
	}
	ServiceInfo struct {
		// ID 驱动实例化后生成的唯一ID，驱动管理服务自动生成。
		// 驱动实例启动后会通过该ID去元数据服务同步设备和更新驱动配置。
		ID     string
		Name   string
		Server RPCConfig
		// ProductList 驱动对应的产品ID列表
		//ProductList []string
		//GwId        string
		//LocalKey    string
		// 跳过激活检查
		Activated bool
	}

	DriverConfig struct {
		Logger      LogConfig
		Clients     map[string]ClientInfo
		Service     ServiceInfo
		CustomParam string
	}

	CloudInstanceConfig struct {
		Logger        CloudInstanceLogConfig
		Clients       map[string]ClientInfo
		Authorization AuthorizationInfo
		Service       ServiceInfo
	}
	AuthorizationInfo struct {
		AK         string
		SK         string
		Regions    string
		ProjectId  string
		InstanceId string
		Endpoint   string
		MqttHost   string
		MqttPort   string
	}

	AppServiceConfig struct {
		Log struct {
			LogLevel string
			LogPath  string
		}
		Tedge struct {
			Host string
			Port int32
		}
		Server struct {
			ID   string
			Name string
			Host string
			Port int32
		}
		//应用私有配置
		CustomConfig map[string]interface{}
	}
)
