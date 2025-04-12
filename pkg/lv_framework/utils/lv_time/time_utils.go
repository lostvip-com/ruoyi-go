package lv_time

import (
	"time"
)

// 获取相差时间
func GetHourDiffer(start_time, end_time string) int64 {
	var hour int64
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", start_time, time.Local)
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", end_time, time.Local)
	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix() //
		hour = diff / 3600
		return hour
	} else {
		return hour
	}
}

func ParseTime(str string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"   // 定义时间格式布局
	t, err := time.Parse(layout, str) // 解析字符串
	if err != nil {
		return t, err
	}
	return t, nil // 返回时间指针
}
