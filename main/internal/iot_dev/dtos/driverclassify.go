package dtos

type DriverClassifyResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type DriverClassifyQueryRequest struct {
	BaseSearchConditionQuery `schema:",inline"`
	Name                     string `schema:"name"`
}
