package lv_sql

type GeneralDelDto struct {
	Id  int   `uri:"id" json:"id" validate:"required"`
	Ids []int `json:"ids"`
}

func (g GeneralDelDto) GetIds() []int {
	ids := make([]int, 0)
	if g.Id != 0 {
		ids = append(ids, g.Id)
	}
	if len(g.Ids) > 0 {
		for _, id := range g.Ids {
			if id > 0 {
				ids = append(ids, id)
			}
		}
	} else {
		if g.Id > 0 {
			ids = append(ids, g.Id)
		}
	}
	if len(ids) <= 0 {
		//方式全部删除
		ids = append(ids, 0)
	}
	return ids
}

type GeneralGetDto struct {
	Id int `uri:"id" json:"id" validate:"required"`
}

//
//func MakeCondition(q interface{}) func(db *lvbatis.DB) *lvbatis.DB {
//	return func(db *lvbatis.DB) *lvbatis.DB {
//		condition := &GormCondition{
//			GormPublic: GormPublic{},
//			Join:       make([]*GormJoin, 0),
//		}
//		ResolveSearchQuery("mysql", q, condition)
//		for _, join := range condition.Join {
//			if join == nil {
//				continue
//			}
//			db = db.Joins(join.JoinOn)
//			for k, v := range join.Where {
//				db = db.Where(k, v...)
//			}
//			for k, v := range join.Or {
//				db = db.Or(k, v...)
//			}
//			for _, o := range join.Order {
//				db = db.Order(o)
//			}
//		}
//		for k, v := range condition.Where {
//			db = db.Where(k, v...)
//		}
//		for k, v := range condition.Or {
//			db = db.Or(k, v...)
//		}
//		for _, o := range condition.Order {
//			db = db.Order(o)
//		}
//		return db
//	}
//}
//
//func Paginate(pageSize, pageIndex int) func(db *lvbatis.DB) *lvbatis.DB {
//	return func(db *lvbatis.DB) *lvbatis.DB {
//		offset := (pageIndex - 1) * pageSize
//		if offset < 0 {
//			offset = 0
//		}
//		return db.Offset(offset).Limit(pageSize)
//	}
//}
