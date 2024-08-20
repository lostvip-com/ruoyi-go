package dtos

const (
	PageDefault        = 1
	PageSizeDefault    = 10
	PageSizeMaxDefault = 1000
)

func CorrectionPageParam(query *BaseSearchConditionQuery) {
	if query.Page <= 0 {
		query.Page = PageDefault
	}

	if query.PageSize >= PageSizeMaxDefault {
		query.PageSize = PageSizeMaxDefault
	} else if query.PageSize <= 0 {
		query.PageSize = PageSizeDefault
	}
}
