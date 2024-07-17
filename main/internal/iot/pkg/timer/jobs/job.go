package jobs

type (
	// RuntimeJobStu 所有任务全部使用一个数据结构 通过time type区分任务类型
	RuntimeJobStu struct {
		JobID       string
		JobName     string
		Description string
		Status      string
		TimeData    TimeData
		JobData     JobData
		Runtimes    int64
	}

	// TimeData 定时表达式类型
	TimeData struct {
		//Type       uint8       `json:"type"`
		Expression string `json:"expression"` // cronExp or repeatExp or sampleExp
	}

	// CronExp crontab表达式
	//0 1 * * * 每天1点
	// 03:04:05 1,2,3,5,6,7 => 4 3 * * 0-2,4-6
	CronExp struct {
		CronTab string `json:"cronTab"`
	}

	// JobData 任务类型
	JobData struct {
		//ActionType uint8       `json:"actionType"` // 1: 场景定时， 2: 设备定时
		ActionData []DeviceMeta `json:"actionData"` // 场景或设备定时 deviceMeta or sceneMeta
	}

	// DeviceMeta 设备定时元数据
	DeviceMeta struct {
		ProductId   string `json:"productId"`
		ProductName string `json:"productName"`
		DeviceId    string `json:"deviceId"`
		DeviceName  string `json:"deviceName"`
		Code        string `json:"code"`
		DateType    string `json:"dateType"`
		Value       string `json:"value"`
	}
)
