package models

type Timestamps struct {
	Created  int64 `json:"created" gorm:"comment:创建时间"`  // Created is a timestamp indicating when the entity was created.
	Modified int64 `json:"modified" gorm:"comment:更新时间"` // Modified is a timestamp indicating when the entity was last modified.
}
