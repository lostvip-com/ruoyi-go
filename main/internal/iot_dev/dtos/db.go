package dtos

type Configuration struct {
	DbType       string
	Host         string
	Port         string
	Timeout      int
	DatabaseName string
	Username     string
	Password     string
	BatchSize    int
	Dsn          string
	// 添加sqlite数据库存储地址
	DataSource string
	// 添加tqlite集群地址
	Cluster []string
}
