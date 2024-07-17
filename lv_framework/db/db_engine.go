package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lostvip-com/lv_framework/db/xorm"
	"github.com/lostvip-com/lv_framework/logme"
	"github.com/lostvip-com/lv_framework/lv_global"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strings"
	"sync"
	"xorm.io/core"
)

type dbEngine struct {
	masterXorm *xorm.Engine //主数据库
	//sqlX
	gormMaster *gorm.DB
	gormSlave  *gorm.DB
}

var (
	instance *dbEngine
	once     sync.Once
)

func init() {
	fmt.Println("-----正在初始化数据库连接-------")
}

// 获取操作实例 如果传入slave 并且成功配置了slave 返回slave orm引擎 否则返回master orm引擎
func (db *dbEngine) GetOrm(dbType string) *gorm.DB {
	if strings.EqualFold(dbType, "slave") {
		return GetSlaveGorm() //随机选取一个
	} else {
		return GetMasterGorm()
	}
	return db.gormMaster
}
func GetMasterGorm() *gorm.DB {
	master := GetInstance().gormMaster
	if master == nil {
		var config = lv_global.Config()
		driverName := config.GetDriver()
		master = createGorm(driverName, config.GetMaster())
		GetInstance().gormMaster = master
	}
	if lv_global.Config().IsDebug() {
		master = master.Debug()
	}
	return master
}

/**
 * 目前只实现一个数据源
 */
func GetSlaveGorm() *gorm.DB {
	slave := GetInstance().gormSlave
	if slave == nil {
		var config = lv_global.Config()
		driverName := config.GetDriver()
		slave = createGorm(driverName, config.GetSlave())
		GetInstance().gormMaster = slave
	}
	return slave
}

func createGorm(driverName, url string) *gorm.DB {
	if !strings.Contains(url, "?") {
		url = url + "?"
	}
	params := "parseTime=true"
	if !strings.Contains(url, params) { //自动解析时间类型到time.Time!!
		if strings.HasSuffix(url, "?") {
			url = url + params
		} else {
			url = url + "&" + params
		}
	}
	params = "charset=utf8mb4"
	if !strings.Contains(url, params) {
		if !strings.Contains(url, params) {
			if strings.HasSuffix(url, "?") {
				url = url + params
			} else {
				url = url + "&" + params
			}
		}
	}
	var dialector gorm.Dialector
	if "mysql" == driverName {
		dialector = mysql.Open(url)
	} else {
		dialector = sqlite.Open(url)
	}
	gormDB, err := gorm.Open(dialector, &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		panic("连接数据库失败" + err.Error())
	}
	sqlDB, err := gormDB.DB() //dr
	if err != nil {
		panic("连接数据库失败")
	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(50)
	logme.Info("########### db conn 初始化成功！ #################")
	return gormDB
}

// /////////////////////////////////////////////////////////////////////////////////////////
// xorm
// ////////////////////////////////////////////////////////////////////////////////////////
func createXormEngine(driverName string, url string, debug bool) *xorm.Engine {
	if !strings.Contains(url, "?") {
		url = url + "?"
	}
	params := "parseTime=true"
	if !strings.Contains(url, params) {
		if strings.HasSuffix(url, "?") {
			url = url + params
		} else {
			url = url + "&" + params
		}
	}
	params = "charset=utf8mb4"
	if !strings.Contains(url, params) {
		if !strings.Contains(url, params) {
			if strings.HasSuffix(url, "?") {
				url = url + params
			} else {
				url = url + "&" + params
			}
		}
	}
	engine, err := xorm.NewEngine(driverName, url)
	if err != nil {
		panic("数据库连接错误:%v" + err.Error())
		return engine
	}
	err = engine.Ping()
	if err != nil {
		panic("数据库连接错误:%v" + err.Error())
		return engine
	}
	if debug {
		engine.ShowSQL(true)
		engine.Logger().SetLevel(core.LOG_DEBUG)
	}

	return engine
}

// 初始化数据操作 driver为数据库类型
func GetInstance() *dbEngine {
	once.Do(func() {
		//driverName := core.SQLITE
		var db dbEngine
		var config = lv_global.Config()
		driverName := config.GetDriver()
		//没有配置从数据库
		if config.GetMaster() != "" {
			//xorm (即将移除xorm,目前作为过渡)
			db.masterXorm = createXormEngine(driverName, config.GetMaster(), false)
		}
		instance = &db
	})
	return instance
}

// 获取操作实例 如果传入slave 并且成功配置了slave 返回slave orm引擎 否则返回master orm引擎
func (db *dbEngine) Engine(dbType ...string) *xorm.Engine {
	if dbType != nil && len(dbType) > 0 {
		if strings.EqualFold(dbType[0], "slave") {
			panic("暂不支持！")
			return nil //随机选取一个
		}
	}

	if db == nil {
		panic("\n ------------>错误信息：\n无法链接到数据库!!!! 检查相关配置，如：\n ------------>masterXorm:\n " + lv_global.Config().GetMaster() + "\n ------------>slave:\n " + lv_global.Config().GetSlave())
	}
	return db.masterXorm
}
