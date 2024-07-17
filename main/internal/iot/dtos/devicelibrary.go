package dtos

import models "main/internal/iot/model"

type DeviceLibrary struct {
	Id             string
	Name           string // 容器名/软件名
	Alias          string
	Description    string
	Protocol       string
	Version        string
	LibFile        string
	ConfigFile     string
	Config         string
	DockerConfigId string
	DockerRepoName string
	DockerImageId  string
	//SupportVersions []DeviceLibrarySupportVersion
	IsInternal    bool
	OperateStatus string // 下载状态
}

func DeviceLibraryFromModel(d models.DeviceLibrary) DeviceLibrary {
	return DeviceLibrary{
		Id:             d.Id,
		Name:           d.Name,
		Description:    d.Description,
		Protocol:       d.Protocol,
		Version:        d.Version,
		DockerConfigId: d.DockerConfigId,
		DockerRepoName: d.DockerRepoName,
		DockerImageId:  d.DockerImageId,
	}
}

type DeviceLibraryAddRequest struct {
	Id             string `json:"id,omitempty"`
	Name           string `json:"name" binding:"required"`
	Description    string `json:"description,omitempty"`
	Protocol       string `json:"protocol"`
	Version        string `json:"version" binding:"required"`
	ContainerName  string `json:"container_name" binding:"required"`
	DockerConfigId string `json:"docker_config_id" binding:"required"`
	DockerRepoName string `json:"docker_repo_name" binding:"required"`
	Language       string `json:"language"` //代码语言
	SupportVersion struct {
		IsDefault          bool   `json:"is_default"`
		DockerParamsSwitch bool   `json:"docker_params_switch"`
		DockerParams       string `json:"docker_params"`
		ExpertMode         bool   `json:"expert_mode"`
		ExpertModeContent  string `json:"expert_mode_content"`
		ConfigJson         string `json:"config_json"`
	} `json:"support_version"`
}

func FromDeviceLibraryRpcToModel(p *DeviceLibraryAddRequest) models.DeviceLibrary {
	dl := models.DeviceLibrary{
		Id:             p.Id,
		Name:           p.Name,
		Description:    p.Description,
		Protocol:       p.Protocol,
		Version:        p.Version,
		ContainerName:  p.ContainerName,
		DockerRepoName: p.DockerRepoName,
		DockerConfigId: p.DockerConfigId,
		Language:       p.Language,
	}
	return dl
}

type DeviceLibrarySearchQueryRequest struct {
	BaseSearchConditionQuery `schema:",inline"`
	DockerConfigId           string `form:"docker_config_id" json:"docker_config_id"`
	IsInternal               string `form:"is_internal" json:"is_internal"`
	ClassifyId               int    `form:"classify_id" json:"classify_id"`
	DockerRepoName           string `form:"docker_repo_name" json:"docker_repo_name"`
	NameAliasLike            string `form:"name_alias_like" json:"name_alias_like"`
	DownloadStatus           string `form:"download_status" json:"download_status"`
	DriverType               int    `form:"driver_type" json:"driver_type" binding:"omitempty,oneof=1 2"` // 驱动库类型，1：驱动，2：三方应用
	NoInIds                  string // 约定，没有from的为 内置查询条件
	ImageIds                 string // 内置条件
	NoInImageIds             string // 内置条件
}

type DeviceLibraryResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	//Alias           string                              `json:"alias"`
	Description string `json:"description"`
	//Manufacturer    string                              `json:"manufacturer"`
	Protocol      string `json:"protocol"`
	Version       string `json:"version"`
	ContainerName string `json:"container_name"`
	//ConfigFile      string                              `json:"config_file"`
	DockerConfigId  string                              `json:"docker_config_id"`
	DockerRepoName  string                              `json:"docker_repo_name"`
	OperateStatus   string                              `json:"operate_status"`
	IsInternal      bool                                `json:"is_internal"`
	Manual          string                              `json:"manual"`
	Icon            string                              `json:"icon"`
	ClassifyId      int                                 `json:"classify_id"`
	Created         int64                               `json:"created"`
	Language        string                              `json:"language"`
	SupportVersions []DeviceLibrarySupportVersionSimple `json:"support_versions"` // 用于前端展示可供下载/更新的版本号 key:value == 版本号:配置文件
}

func DeviceLibraryResponseFromModel(dl models.DeviceLibrary) DeviceLibraryResponse {
	// 如果docker镜像id为空，那么返回给前端的版本为 `-`
	if dl.DockerImageId == "" && !dl.IsInternal {
		dl.Version = "-"
	}
	return DeviceLibraryResponse{
		Id:            dl.Id,
		Name:          dl.Name,
		Description:   dl.Description,
		Protocol:      dl.Protocol,
		Version:       dl.Version,
		ContainerName: dl.ContainerName,
		//ConfigFile:      dl.ConfigFile,
		IsInternal:      dl.IsInternal,
		DockerConfigId:  dl.DockerConfigId,
		DockerRepoName:  dl.DockerRepoName,
		OperateStatus:   dl.OperateStatus,
		Icon:            dl.Icon,
		Manual:          dl.Manual,
		ClassifyId:      dl.ClassifyId,
		Created:         dl.Created,
		Language:        dl.Language,
		SupportVersions: DeviceLibrarySupportVersionSimpleFromModel(dl.SupportVersions),
	}
}

type DeviceLibraryUpgradeRequest struct {
	Id      string `json:"id" binding:"required"`
	Version string `json:"version" binding:"required"`
}

type DeviceLibraryUpgradeResponse struct {
	Id            string `json:"id"`
	Version       string `json:"version"`
	OperateStatus string `json:"operate_status"`
}

func GetLibrarySimpleBaseConfig() string {
	return `{"deviceServer": {},"deviceProtocols": {},"deviceDpAttrs": []}`
}

type UpdateDeviceLibrary struct {
	Id             string  `json:"id" binding:"required"`
	Name           *string `json:"name"`
	Description    *string `json:"description"`
	DockerConfigId *string `json:"docker_config_id"`
	Protocol       *string `json:"protocol"`
	Language       *string `json:"language"`
	Manual         *string `json:"manual"`
	Icon           *string `json:"icon"`
}

func ReplaceDeviceLibraryModelFieldsWithDTO(deviceLibrary *models.DeviceLibrary, patch UpdateDeviceLibrary) {
	if patch.Name != nil {
		deviceLibrary.Name = *patch.Name
	}

	if patch.DockerConfigId != nil {
		deviceLibrary.DockerConfigId = *patch.DockerConfigId
	}

	if patch.Description != nil {
		deviceLibrary.Description = *patch.Description
	}

	if patch.Protocol != nil {
		deviceLibrary.Protocol = *patch.Protocol
	}

	if patch.Language != nil {
		deviceLibrary.Language = *patch.Language
	}

	if patch.Manual != nil {
		deviceLibrary.Manual = *patch.Manual
	}
	if patch.Icon != nil {
		deviceLibrary.Icon = *patch.Icon
	}
}
