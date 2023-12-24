package user_role

import (
	"lostvip.com/db"
)

// Fill with you ideas below.
// 批量删除
func DeleteBatch(ids ...int64) (int64, error) {
	return db.GetInstance().Engine().Table(TableName()).In("user_id", ids).Delete(new(Entity))
}
