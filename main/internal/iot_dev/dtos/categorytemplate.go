package dtos

type CategoryTemplateRequest struct {
	BaseSearchConditionQuery `schema:",inline"`
	CategoryName             string `schema:"category_name"`
	Scene                    string `schema:"scene"`
}

type CategoryTemplateSyncRequest struct {
	VersionName string `json:"version_name"`
}

type ThingModelTemplateSyncRequest struct {
	VersionName string `json:"version_name"`
}

type CategoryTemplateResponse struct {
	Id           string `json:"id"`
	CategoryName string `json:"category_name"` //品类名称
	CategoryKey  string `json:"category_key"`
	Scene        string `json:"scene"` //所属场景
}

type CosCategoryTemplateResponse struct {
	CategoryName string `json:"category_name"`
	CategoryKey  string `json:"category_key"`
	Scene        string `json:"scene"`
}

type ThingModelTemplateRequest struct {
	BaseSearchConditionQuery `schema:",inline"`
	CategoryKey              string `schema:"category_key"`
	CategoryName             string `schema:"category_name"`
}

type ThingModelTemplateResponse struct {
	Id             string `json:"id"`
	CategoryName   string `json:"category_name"` //品类名称
	CategoryKey    string `json:"category_key"`
	ThingModelJSON string `json:"thing_model_json"`
	//models.Properties
	Properties interface{} `json:"properties"`
	Events     interface{} `json:"events"`
	Actions    interface{} `json:"actions"`
}

type CosThingModelTemplateResponse struct {
	CategoryName   string `json:"category_name"`
	CategoryKey    string `json:"category_key"`
	ThingModelJSON string `json:"thing_model_json"`
}
