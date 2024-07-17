package models

// 驱动镜像配置表，alias MirrorConfig
type DockerConfig struct {
	Timestamps `gorm:"embedded"`
	Id         string `gorm:"id;primaryKey;not null;type:string;size:255;comment:主键"`
	Address    string `gorm:"type:string;size:255;comment:地址"`
	Account    string `gorm:"type:string;size:255;comment:用户名"`
	Password   string `gorm:"type:string;size:255;comment:密码"`
	SaltKey    string `gorm:"type:string;size:255;comment:盐值"`
}

func (table *DockerConfig) TableName() string {
	return "docker_config"
}

func (table *DockerConfig) Get() interface{} {
	return *table
}
