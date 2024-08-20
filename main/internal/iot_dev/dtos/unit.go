package dtos

type UnitRequest struct {
	BaseSearchConditionQuery `schema:",inline"`
	UnitName                 string `schema:"unitName" json:"unitName"`
}

type UnitResponse struct {
	Id       string `json:"id"`
	Symbol   string `json:"symbol"`
	UnitName string `json:"unit_name"`
}

type CosUnitTemplateResponse struct {
	UnitName string `json:"Name"`
	Symbol   string `json:"Symbol"`
}

type UnitTemplateSyncRequest struct {
	VersionName string `json:"version_name"`
}
