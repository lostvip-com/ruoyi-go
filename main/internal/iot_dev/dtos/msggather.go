package dtos

type MsgGatherSearchQueryRequest struct {
	BaseSearchConditionQuery `schema:",inline"`
	Date                     []string `schema:"date"`
}
