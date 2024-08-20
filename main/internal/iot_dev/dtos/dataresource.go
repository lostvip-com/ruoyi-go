package dtos

import (
	models "main/internal/iot_dev/model"
	"main/internal/iot_dev/pkg/constants"
)

type DataResourceSearchQueryRequest struct {
	BaseSearchConditionQuery `schema:",inline"`
	Type                     string `schema:"type,omitempty"`
	Health                   string `schema:"health,omitempty"`
}

type DataResourceInfo struct {
	Name   string                 `json:"name"`
	Type   string                 `json:"type"`
	Option map[string]interface{} `json:"option"`
}

type AddDataResourceReq struct {
	Name   string                 `json:"name"`
	Type   string                 `json:"type"`
	Option map[string]interface{} `json:"option"`
}

type UpdateDataResource struct {
	Id     string                  `json:"id"`
	Name   *string                 `json:"name"`
	Type   *string                 `json:"type"`
	Option *map[string]interface{} `json:"option"`
}

func ReplaceDataResourceModelFields(ds *models.DataResource, patch UpdateDataResource) {
	if patch.Name != nil {
		ds.Name = *patch.Name
	}
	if patch.Type != nil {
		ds.Type = constants.DataResourceType(*patch.Type)
	}
	if patch.Option != nil {
		ds.Option = *patch.Option
		ds.Option["sendSingle"] = true
	}
	ds.Health = false

}
