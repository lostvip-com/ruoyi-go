package lv_sql

//
//import (
//	"fmt"
//	"testing"
//	"time"
//)
//
//type ApplicationQuery struct {
//	Id       string    `lv_sql:"type:icontains;column:id;table:receipt" form:"id"`
//	Domain   string    `lv_sql:"type:icontains;column:domain;table:receipt" form:"domain"`
//	Version  string    `lv_sql:"type:exact;column:version;table:receipt" form:"version"`
//	Status   []int     `lv_sql:"type:in;column:status;table:receipt" form:"status"`
//	Start    time.Time `lv_sql:"type:gte;column:created_at;table:receipt" form:"start"`
//	End      time.Time `lv_sql:"type:lte;column:created_at;table:receipt" form:"end"`
//	TestJoin `lv_sql:"type:left;on:id:receipt_id;table:receipt_goods;join:receipts"`
//	NotNeed  string `lv_sql:"-"`
//	ApplicationOrder
//}
//
//type ApplicationOrder struct {
//	IdOrder string `lv_sql:"type:order;column:id;table:receipt" form:"id_order"`
//}
//
//type TestJoin struct {
//	PaymentAccount string `lv_sql:"type:icontains;column:payment_account;table:receipts" form:"payment_account"`
//}
//
//func TestResolveSearchQuery(t *testing.T) {
//	// Only pass t into top-level Convey calls
//	Convey("Given some integer with a starting value", t, func() {
//		d := ApplicationQuery{
//			Id:               "aaa",
//			Domain:           "bbb",
//			Version:          "ccc",
//			Status:           []int{1, 2},
//			Start:            time.Now().Add(-8 * time.Hour),
//			End:              time.Now(),
//			ApplicationOrder: ApplicationOrder{IdOrder: "desc"},
//			TestJoin:         TestJoin{PaymentAccount: "1212"},
//		}
//		condition := &GormCondition{
//			GormPublic: GormPublic{},
//			Join:       make([]*GormJoin, 0),
//		}
//		ResolveSearchQuery("mysql", d, condition)
//		fmt.Println(condition)
//
//		var data model_cmn.Settings
//		var list []model_cmn.Settings
//		database.GormDB.Model(&data).
//			Scopes(
//				MakeCondition(data),
//			).Find(&list)
//	})
//}
