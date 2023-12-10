package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"lostvip.com/conf"
	"lostvip.com/db/sqlx"
	"lostvip.com/logme"
	"lostvip.com/utils/lv_sql"
	"strings"
	"sync"
	"xorm.io/core"
)

type dbEngine struct {
	master *xorm.Engine //主数据库
	slave  *xorm.Engine //从数据库
	//sqlX
	sqlxMaster *sqlx.DB
	sqlxSlave  *sqlx.DB
}

var (
	instance *dbEngine
	once     sync.Once
)

func init() {
	//这里设置sqlx中，在使用interface{}类型传sql参数时，
	//占位符的格式默认是全部小写：这里改为首字母大写，
	//如： and post_name like :PostName
	sqlx.NameMapper = lv_sql.ToTitle
}

func createSqlx(dsn string) *sqlx.DB {
	// 连接到数据库并使用ping进行验证。
	// 也可以使用 MustConnect MustConnect连接到数据库，并在出现错误时恐慌 panic。
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic("connect DB failed, err:%v\n" + err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic("connect DB failed, err:%v\n" + err.Error())
	}
	db.SetMaxOpenConns(100) // 设置数据库的最大打开连接数。
	db.SetMaxIdleConns(1)   // 设置空闲连接池中的最大连接数。
	logme.Info("########### sqlx 初始化成功！ #################")
	return db
}

/**
 * 获取主库的实例
 */
func GetInstanceMaster() *xorm.Engine {
	return GetInstance().Engine()
}

func GetMasterSqlX() *sqlx.DB {
	return GetInstance().sqlxMaster
}

func GetSlaveSqlX() *sqlx.DB {
	return GetInstance().sqlxSlave
}

/**
 * 获取从库的实例
 */
func GetInstanceSlave() *xorm.Engine {
	return GetInstance().Engine("slave")
}

/**
 * 单例方法
 */
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
		if config.GetMaster() != "" {
			//xorm
			db.master = createXormEngine(driverName, config.GetMaster(), config.IsDebug())
			//sqlx
			db.sqlxMaster = createSqlx(config.GetMaster())
		}
		if config.GetSlave() != "" {
			//xorm
			db.slave = createXormEngine(driverName, config.GetSlave(), config.IsDebug())
			//sqlx
			db.sqlxSlave = createSqlx(config.GetSlave())
		}
		instance = &db
	})
	return instance
}

func createXormEngine(driverName string, url string, debug bool) *xorm.Engine {
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

// 获取操作实例 如果传入slave 并且成功配置了slave 返回slave orm引擎 否则返回master orm引擎
func (db *dbEngine) Engine(dbType ...string) *xorm.Engine {
	if dbType != nil && len(dbType) > 0 {
		if strings.EqualFold(dbType[0], "slave") {
			if db.slave != nil {
				return db.slave //随机选取一个
			}
		}
	}
	if db == nil {
		panic("\n ------------>错误信息：\n无法链接到数据库!!!! 检查相关配置，如：\n ------------>master:\n " + conf.Config().GetMaster() + "\n ------------>slave:\n " + conf.Config().GetSlave())
	}
	return db.master
}
