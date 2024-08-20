package lv_drivers

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// DbConn mysql配置结构体
type DbConn struct {
	User   string
	Pwd    string
	Host   string
	Port   string
	DbName string
	//
	DbType string
	gormDB *gorm.DB
}

/**
 * 测试并关闭数据库连接
 */
func (e *DbConn) TestConnInstance(User, Pwd, Host, Port, DbName string) {
	myDbConn := &DbConn{User: User, Pwd: Pwd, Host: Host, Port: Port, DbName: DbName}
	g, err := myDbConn.Setup()
	if err != nil {
		panic(err)
	}
	sqlDB, err := g.DB()
	sqlDB.Close()
}
func (e *DbConn) GetGormDB() *gorm.DB {
	if e.gormDB == nil {
		e.gormDB, _ = e.Setup()
	}
	return e.gormDB
}

// Setup 配置步骤
//
//root:root!@tcp(60.205.205.242:13307)/main.com?charset=utf8&parseTime=True&loc=Local&timeout=1000ms
func (e *DbConn) Setup() (*gorm.DB, error) {
	var err error
	url := "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=5000ms"
	url = fmt.Sprintf(url, e.User, e.Pwd, e.Host, e.Port, e.DbName)
	dialector := mysql.Open(url)
	e.gormDB, err = gorm.Open(dialector, &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		panic("连接数据库失败" + err.Error())
	}
	sqlDB, err := e.gormDB.DB() //dr
	if err != nil {
		panic("连接数据库失败")
	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(50)
	//e.GormSqlite.LogMode(true) // ====================打印sql
	return e.gormDB, err
}

//
//func (e *DbConn) CloseConn() {
//	sqlDB, err := e.gormDB.DB() //dr
//	if err == nil {
//		sqlDB.Close()
//	}
//	myDbConn = nil
//	Logger.Info(" close success !")
//	return
//}
