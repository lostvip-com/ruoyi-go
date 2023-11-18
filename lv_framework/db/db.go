package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"lostvip.com/conf"
	"strings"
	"sync"
	"xorm.io/core"
)

type dbEngine struct {
	master *xorm.Engine //主数据库
	slave  *xorm.Engine //从数据库
}

var (
	instance *dbEngine
	once     sync.Once
)

func GetInstance() *dbEngine {
	if instance == nil {
		instance = Instance()
	}
	return instance
}

// 初始化数据操作 driver为数据库类型
func Instance() *dbEngine {
	once.Do(func() {
		//driverName := core.SQLITE
		var db dbEngine
		var config = conf.Config()
		driverName := config.GetDriver()
		//没有配置从数据库
		if len(config.GetSlave()) == 0 {
			engine, err := xorm.NewEngine(driverName, config.GetMaster())
			if err != nil {
				fmt.Printf("数据库连接错误:%v", err.Error())
				return
			}
			err = engine.Ping()
			if err != nil {
				fmt.Printf("数据库连接错误:%v", err.Error())
				return
			}
			if config.IsDebug() {
				engine.ShowSQL(true)
				engine.Logger().SetLevel(core.LOG_DEBUG)
			}
			db.master = engine
			instance = &db
		} else {
			master, err := xorm.NewEngine(driverName, config.GetMaster())
			if err != nil {
				return
			}
			if config.IsDebug() {
				master.ShowSQL(true)
				master.Logger().SetLevel(core.LOG_DEBUG)
			}

			slave, err := xorm.NewEngine(driverName, config.GetSlave())
			if err != nil {
				return
			}
			slaves := []*xorm.Engine{slave}
			group, err := xorm.NewEngineGroup(master, slaves)
			if config.IsDebug() {
				group.ShowSQL(true)
				group.Logger().SetLevel(core.LOG_DEBUG)
			}
			db.master = group.Master()
			db.slave = group.Slave() //随机选一个
			instance = &db
		}
	})
	return instance
}

// 获取操作实例 如果传入slave 并且成功配置了slave 返回slave orm引擎 否则返回master orm引擎
func (db *dbEngine) Engine(dbType ...string) *xorm.Engine {
	if dbType != nil && len(dbType) > 0 {
		if strings.EqualFold(dbType[0], "slave") {
			if db.slave != nil {
				return db.slave
			}
		}
	}
	return db.master
}
