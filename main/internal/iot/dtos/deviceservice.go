package dtos

import (
	"github.com/winc-link/hummingbird/internal/models"
	"github.com/winc-link/hummingbird/internal/pkg/constants"

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

func DeviceServiceFromModel(ds models.DeviceService) DeviceService {
	var dto DeviceService
	dto.Id = ds.Id
	dto.Name = ds.Name
	dto.BaseAddress = ds.BaseAddress
	dto.DeviceLibraryId = ds.DeviceLibraryId
	dto.Config = ds.Config
	dto.DockerContainerId = ds.DockerContainerId
	dto.ExpertMode = ds.ExpertMode
	dto.ExpertModeContent = ds.ExpertModeContent
	dto.DockerParamsSwitch = ds.DockerParamsSwitch
	dto.DockerParams = ds.DockerParams
	dto.ContainerName = ds.ContainerName
	return dto
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
	Platform           constants.IotPlatform   `json:"platform"`
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

func ReplaceDeviceServiceModelFieldsWithDTO(ds *models.DeviceService, patch DeviceServiceUpdateRequest) {
	if patch.Config != nil {
		ds.Config = *patch.Config
	}
	if patch.DeviceLibraryId != nil {
		ds.DeviceLibraryId = *patch.DeviceLibraryId
	}
	if patch.ExpertMode != nil {
		ds.ExpertMode = *patch.ExpertMode
	}
	if patch.ExpertModeContent != nil {
		ds.ExpertModeContent = *patch.ExpertModeContent
	}
	if patch.DockerParamsSwitch != nil {
		ds.DockerParamsSwitch = *patch.DockerParamsSwitch
	}
	if patch.DockerParams != nil {
		ds.DockerParams = *patch.DockerParams
	}
	if patch.Platform != "" {
		ds.Platform = patch.Platform
	}
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

func DeviceServiceResponseFromModel(ds models.DeviceService, dl models.DeviceLibrary) DeviceServiceResponse {
	return DeviceServiceResponse{
		Id:   ds.Id,
		Name: ds.Name,
		//Version: DeviceLibraryResponseFromModel(dl).Version,
		DeviceLibrary:      DeviceLibraryResponseFromModel(dl),
		RunStatus:          ds.RunStatus,
		Config:             ds.Config,
		ExpertMode:         ds.ExpertMode,
		ExpertModeContent:  ds.ExpertModeContent,
		DockerParamsSwitch: ds.DockerParamsSwitch,
		DockerParams:       ds.DockerParams,
		ImageExist:         ds.ImageExist,
		CreateAt:           ds.Created,
		Platform:           string(ds.Platform),
	}
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
