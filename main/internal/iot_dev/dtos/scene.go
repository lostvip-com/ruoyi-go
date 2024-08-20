package dtos

type SceneAddRequest struct {
	Name        string `json:"name"`        //名字
	Description string `json:"description"` //描述
}

type SceneUpdateRequest struct {
	Id         string      `json:"id"`
	Conditions []Condition `json:"conditions"`
	Actions    []Action    `json:"actions"`
}

type Condition struct {
	ConditionType string            `json:"condition_type"`
	Option        map[string]string `json:"option"`
}

type Action struct {
	ProductID   string `json:"product_id"`
	ProductName string `json:"product_name"`
	DeviceID    string `json:"device_id"`
	DeviceName  string `json:"device_name"`
	Code        string `json:"code"`
	DataType    string `json:"data_type"`
	Value       string `json:"value"`
}

type SceneSearchQueryRequest struct {
	BaseSearchConditionQuery `schema:",inline"`
	Name                     string `json:"name"`
	Status                   string `json:"status"`
}

type SceneLogSearchQueryRequest struct {
	BaseSearchConditionQuery `schema:",inline"`
	StartAt                  int64  `schema:"start_time"`
	EndAt                    int64  `schema:"end_time"`
	SceneId                  string `json:"scene_id"`
}
