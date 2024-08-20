// ==========================================================================
// LV自动生成数据库操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2024-07-19 17:09:35 +0800 CST
// 生成人：lv
// ==========================================================================
package model

import (
	"common/model"
	"github.com/lostvip-com/lv_framework/lv_db"
	"time"
)

// IotDevice IotDevice
type IotDevice struct {
	Id        int64 `gorm:"type:bigint(20);primary_key;auto_increment;;" json:"id"`
	DeptId    int64 `gorm:"type:bigint(20);comment:上级单位ID;" json:"deptId"`
	ProductId int64 `gorm:"type:bigint(20);comment:产品ID;" json:"productId"`

	GatewayId       int64      `gorm:"type:bigint(20);comment:所属网关ID，只对网关子设备有效;" json:"gatewayId"`
	DriveInstanceId string     `gorm:"type:varchar(32);comment:驱动实例ID;" json:"driveInstanceId"`
	NodeType        string     `gorm:"type:varchar(32);comment:节点类型3设备2网关;" json:"nodeType"`
	Name            string     `gorm:"type:varchar(32);comment:名称;" json:"name"`
	Sn              string     `gorm:"type:varchar(32);comment:序列号;" json:"sn"`
	DevNo           int16      `gorm:"type:int;comment:从机地址;" json:"devNo"`
	Status          string     `gorm:"type:int;default:0;comment:设备状态;-1 离线 ,0工程态，1 上线 " json:"status"`
	Secret          string     `gorm:"type:varchar(32);comment:密钥;" json:"secret"`
	Description     string     `gorm:"type:text;comment:描述;" json:"description"`
	InstallLocation string     `gorm:"type:varchar(32);comment:安装地址;" json:"installLocation"`
	LastSyncTime    *time.Time `gorm:"type:datetime;comment:最后一次同步时间;" json:"lastSyncTime" time_format:"2006-01-02 15:04:05"`
	LastOnlineTime  *time.Time `gorm:"type:datetime;comment:最后一次在线时间;" json:"lastOnlineTime" time_format:"2006-01-02 15:04:05"`
	model.BaseModel
	// 非映射字段
	ProductName  string `gorm:"-" json:"productName"`
	NodeTypeName string `gorm:"-" json:"nodeTypeName"`
}

func (e *IotDevice) TableName() string {
	return "iot_device"
}

// 增
func (e *IotDevice) Save() error {
	return lv_db.GetMasterGorm().Save(e).Error
}

// 查
func (e *IotDevice) FindById(deviceId int64) (*IotDevice, error) {
	err := lv_db.GetMasterGorm().Take(e, deviceId).Error
	return e, err
}

// 查第一条
func (e *IotDevice) FindOne() error {
	tb := lv_db.GetMasterGorm().Table(e.TableName())

	if e.Name != "" {
		tb = tb.Where("name like ?", "%"+e.Name+"%")
	}
	if e.ProductId != 0 {
		tb = tb.Where("product_id=?", e.ProductId)
	}

	err := tb.First(e).Error
	return err
}

// 改
func (e *IotDevice) Updates() error {
	return lv_db.GetMasterGorm().Updates(e).Error
}

// 删
func (e *IotDevice) Delete() error {
	return lv_db.GetMasterGorm().Delete(e).Error
}
