package lv_db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lostvip-com/lv_framework/lv_global"
	"github.com/lostvip-com/lv_framework/lv_log"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strings"
	"sync"
)

type dbEngine struct {
	gormMap map[string]*gorm.DB
}

var (
	instance *dbEngine
	once     sync.Once
)

func init() {
	fmt.Println("-----正在初始化数据库连接-------")
}

// GetInstance 初始化数据操作 driver为数据库类型
func GetInstance() *dbEngine {
	once.Do(func() {
		instance = new(dbEngine)
		instance.gormMap = make(map[string]*gorm.DB)
	})
	return instance
}

// GetOrm 获取操作实例 如果传入slave 并且成功配置了slave 返回slave orm引擎 否则返回master orm引擎
func (db *dbEngine) GetOrm(dbType string) *gorm.DB {
	gdb := db.gormMap[dbType]
	if gdb == nil {
		return GetMasterGorm()
	}
	return gdb
}
func GetMasterGorm() *gorm.DB {
	master := GetInstance().gormMap["master"]
	if master == nil {
		var config = lv_global.Config()
		driverName := config.GetDriver()
		master = createGorm(driverName, config.GetMaster())
		GetInstance().gormMap["master"] = master
	}
	if lv_global.IsDebug {
		master = master.Debug()
	}
	return master
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

	showSql := lv_global.Config().GetBool("go.datasource.show-sql")
	config := &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}} //表名用单数
	if showSql {
		config.Logger = logger.Default.LogMode(logger.Info)
	}
	gormDB, err := gorm.Open(dialector, config)
	if err != nil {
		panic("连接数据库失败" + err.Error())
	}
	sqlDB, err := gormDB.DB() //dr
	if err != nil {
		panic("连接数据库失败")
	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(50)
	lv_log.Info("########### lv_db conn 初始化成功！ #################")
	return gormDB
}
