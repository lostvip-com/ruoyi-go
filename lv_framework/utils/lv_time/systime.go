package lv_time

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

/**
 * 同步操作系统时间
 */
//func UpdateSystemDate(ntpdateUrl string) error {
//	if ntpdateUrl == "" {
//		return errors.New("请配置时间同步服务器地址")
//	}
//	ntpdateUrl = strings.TrimSpace(ntpdateUrl)
//	//先杀ntpd进程，防止端口123冲突
//	data, err := exec.Command("./stop-ntpd.sh").Output()
//	fmt.Println("杀ntpd: ./stop-ntpd.sh --------->", string(data), "<-----")
//	out, err := lv_os.RunCommand("ntpdate " + ntpdateUrl)
//	fmt.Printf("ntpdate同步时间完成=========%v", out)
//	return err
//}

/**
 * 修改操作系统时区
 */
func UpdateTimeZone(timeZone string) (err error) {
	cmd0 := exec.Command("rm", "-rf", "/etc/localtime")
	out, err := cmd0.Output()
	if err != nil {
		fmt.Println("XXXXXXXXXX rm -rf /etc/localtime  XXXXXX" + string(out) + "error:" + err.Error())
		return err
	}
	timeZone = strings.TrimSpace(timeZone)
	cmdStr := `/usr/share/zoneinfo/` + timeZone
	cmd2 := exec.Command("ln", "-s", cmdStr, "/etc/localtime")
	out, err = cmd2.Output()
	if err != nil {
		err = errors.New("ln -s  XXXXXX " + string(out) + "error:" + err.Error())
		return err
	}
	return err
}

// GetCurrTimeStr  自定义格式
func GetCurrTimeStr(fmt string) string {
	return time.Now().Local().Format(fmt)
}

func GetCurrentTimeStr() string {
	return time.Now().Local().Format("2006-01-02 15:04:05")
}

func GetCurrentDayStr() string {
	return time.Now().Local().Format("2006-01-02")
}

func GetCurrentTime() time.Time {
	return time.Now().Local()
}

func AddMinute(date time.Time, m int) time.Time {
	str := fmt.Sprintf("%dm", m)
	dt, _ := time.ParseDuration(str)
	m1 := date.Add(dt)
	return m1
}

func TimeAddHour(date time.Time, h int) time.Time {
	str := fmt.Sprintf("%dh", h)
	dt, _ := time.ParseDuration(str)
	m1 := date.Add(dt)
	return m1
}

func TimeAddDay(date time.Time, day int) time.Time {
	h := 24 * day
	str := fmt.Sprintf("%dh", h)
	dt, _ := time.ParseDuration(str)
	m1 := date.Add(dt)
	return m1
}

func GetTimeStr(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}

func GetSeconds(datetime string) (int64, error) {
	//日期转化为时间戳
	timeLayout := "2006-01-02 15:04:05"    //转化所需模板
	loc, err := time.LoadLocation("Local") //获取时区
	tmp, err := time.ParseInLocation(timeLayout, datetime, loc)
	timestamp := tmp.Unix() //转化为时间戳 类型是int64
	return timestamp, err
}
