package dtos

type QuickNavigationSearchQueryRequest struct {
	BaseSearchConditionQuery `schema:",inline"`
	Name                     string `schema:"name,omitempty"`
}

type DocsSearchQueryRequest struct {
	BaseSearchConditionQuery `schema:",inline"`
	Name                     string `schema:"name,omitempty"`
}

type CosQuickNavigationTemplateResponse struct {
	Name     string `json:"Name"`
	Sort     int    `json:"Sort"`
	Icon     string `json:"Icon"`
	JumpLink string `json:"JumpLink"`
}

type CosDocTemplateResponse struct {
	Name     string `json:"Name"`
	Sort     int    `json:"Sort"`
	JumpLink string `json:"JumpLink"`
}
