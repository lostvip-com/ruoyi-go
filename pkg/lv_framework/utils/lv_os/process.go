package lv_os

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

///**
//
//IsExeRuning : 程序是否运行，
//strKey:用于结果查找，
//strExeName:校验查找的是否是要寻找的进程
//
// */
//func IsExeRuning(strKey string) bool {
//	buf := bytes.Buffer{}
//	cmd := exec.Command("wmic", "process", "get", "ProcessId,name,executablepath")
//	cmd.Stdout = &buf
//	cmd.Run()
//
//	cmd2 := exec.Command("findstr", strKey)
//	cmd2.Stdin = &buf
//	data, err := cmd2.CombinedOutput()
//	fmt.Println("IsExeRuning strKey: %d --------------- err :%v",strKey,err)
//	if err != nil && err.Error() != "exit status 1" {
//		//XBLogF("ServerMonitor", "IsExeRuning CombinedOutput error, err:%s", err.Error())
//		return false
//	}
//
//	strData := string(data)
//	fmt.Println("IsExeRuning strKey: %d ==================== strData :%s",strKey,strData)
//	if strings.Contains(strData, strKey) {
//		return true
//	} else {
//		return false
//	}
//}
//

func runInWindows(cmd string) (string, error) {
	result, err := exec.Command("cmd", "/c", cmd).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(result)), err
}

func RunCommand(cmd string) (string, error) {
	if runtime.GOOS == "windows" {
		return runInWindows(cmd)
	} else {
		return runInLinux(cmd)
	}
}

func runInLinux(cmd string) (string, error) {
	fmt.Println("Running Linux cmd:" + cmd)
	result, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(result)), err
}

// 根据进程名判断进程是否运行
func CheckProRunning(serverName string) bool {
	a := `ps ux | awk '/` + serverName + `/ && !/awk/ {print $2}'`
	pid, err := RunCommand(a)

	var running = true
	if err != nil || pid == "" {
		running = false
	}
	fmt.Println(serverName, "#################pid: ", pid, "running", running)
	return running
}

// 根据进程名称获取进程ID
func GetPid(serverName string) (string, error) {
	a := `ps ux | awk '/` + serverName + `/ && !/awk/ {print $2}'`
	pid, err := RunCommand(a)
	return pid, err
}
