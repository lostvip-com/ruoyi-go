package dtos

import (
	//"gitlab.com/tedge/edgex/internal/models"
	//devicelibraryProto "gitlab.com/tedge/edgex/proto/devicelibrary"
	//deviceserviceProto "gitlab.com/tedge/edgex/proto/deviceservice"
	"gopkg.in/yaml.v3"
)

type DeviceService struct {
	Id                 string                 `json:"id,omitempty"`
	Name               string                 `json:"name"`
	Created            int64                  `json:"created,omitempty"`
	Modified           int64                  `json:"modified,omitempty"`
	BaseAddress        string                 `json:"baseAddress"`
	DeviceLibraryId    string                 `json:"deviceLibraryId"`
	Config             map[string]interface{} `json:"config"`
	DockerContainerId  string                 `json:"dockerContainerId"`
	ExpertMode         bool                   `json:"isExpertMode"`
	ExpertModeContent  string                 `json:"expertModeContent"`
	DockerParamsSwitch bool                   `json:"dockerParamsSwitch"`
	DockerParams       string                 `json:"dockerParams"`
	ContainerName      string                 `json:"container_name"`
}

// 启动实例时对应的配置
type RunServiceCfg struct {
	ImageRepo          string
	RunConfig          string
	DockerMountDevices []string
	DockerParams       string
	DriverName         string // 驱动名
}

type DeviceServiceAddRequest struct {
	Id                 string                 `json:"id,omitempty" binding:"omitempty,t-special-char"`
	Name               string                 `json:"name"`
	DeviceLibraryId    string                 `json:"deviceLibraryId" binding:"required"`
	Config             map[string]interface{} `json:"config" binding:"required"`
	ExpertMode         bool                   `json:"expertMode"`
	ExpertModeContent  string                 `json:"expertModeContent"`
	DockerParamsSwitch bool                   `json:"dockerParamsSwitch"`
	DockerParams       string                 `json:"dockerParams"`
	DriverType         int                    `json:"driverType" binding:"omitempty,oneof=1 2"` //驱动库类型，1：驱动，2：三方应用
}

type DeviceServiceUpdateRequest struct {
	Id                 string                  `json:"id" binding:"required"`
	DeviceLibraryId    *string                 `json:"deviceLibraryId"`
	Name               *string                 `json:"name"`
	Config             *map[string]interface{} `json:"config"`
	ExpertMode         *bool                   `json:"expertMode"`
	ExpertModeContent  *string                 `json:"expertModeContent"`
	DockerParamsSwitch *bool                   `json:"docker_params_switch"`
	DockerParams       *string                 `json:"docker_params"`
	//IsIgnoreRunStatus  bool
}

type UpdateDeviceServiceRunStatusRequest struct {
	Id        string `json:"id"`
	RunStatus int    `json:"run_status"  binding:"required,oneof=1 2"`
}

type DeviceServiceRunLogRequest struct {
	Id      string `json:"id"`
	Operate int    `json:"operate" binding:"required,oneof=1 2"`
}

type DeviceServiceDeleteRequest struct {
	Id string `json:"id" binding:"required"`
}

type DeviceServiceSearchQueryRequest struct {
	BaseSearchConditionQuery `schema:",inline"`
	ProductId                string `form:"productId"`
	CloudProductId           string `form:"cloudProductId"`
	DeviceLibraryId          string `form:"deviceLibraryId"`  // 驱动库ID
	DeviceLibraryIds         string `form:"deviceLibraryIds"` // 驱动库IDs
	Platform                 string `form:"platform"`
	DriverType               int    `form:"driver_type" binding:"omitempty,oneof=1 2"` //驱动库类型，1：驱动，2：三方应用
}

/************** Response **************/

type DeviceServiceResponse struct {
	Id            string                `json:"id"`
	Name          string                `json:"name"`
	DeviceLibrary DeviceLibraryResponse `json:"deviceLibrary"`
	//Version       string                `json:"version"`
	RunStatus          int         `json:"runStatus"`
	Config             interface{} `json:"config"`
	ExpertMode         bool        `json:"expertMode"`
	ExpertModeContent  string      `json:"expertModeContent"`
	DockerParamsSwitch bool        `json:"dockerParamsSwitch"`
	DockerParams       string      `json:"dockerParams"`
	CreateAt           int64       `json:"create_at"`
	ImageExist         bool        `json:"imageExist"`
	Platform           string      `json:"platform"`
}

func FromYamlStrToMap(yamlStr string) (m map[string]interface{}, err error) {
	err = yaml.Unmarshal([]byte(yamlStr), &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

type UpdateDeviceServiceRunStatusResponse struct {
	Id        string `json:"id"`
	RunStatus int    `json:"run_status"`
}

type UpdateServiceLogLevelConfigRequest struct {
	Id       string `json:"id"` // 驱动或应用ID
	LogLevel int64  `json:"logLevel"`
}
