package utils

import (
    "fmt"
    "net"
    "strconv"
    "strings"
    "time"
)

// CheckNetIface 检查网卡
func CheckNetIface(ethName string) bool {
    return strings.HasPrefix(ethName, "en") || strings.HasPrefix(ethName, "eth")
}

// NetIfaces 获取网卡
func NetIfaces() ([]string, error) {
    interfaces, err := net.Interfaces()
    if err != nil {
        return nil, err
    }
    
    var ifaces []string
    for _, inter := range interfaces {
        if CheckNetIface(inter.Name) {
            ifaces = append(ifaces, inter.Name)
        }
    }
    
    return ifaces, nil
}

func NetMacs() ([]string, error) {
    interfaces, err := net.Interfaces()
    if err != nil {
        return nil, err
    }
    
    var ifaces []string
    for _, inter := range interfaces {
        if CheckNetIface(inter.Name) {
            ifaces = append(ifaces, inter.HardwareAddr.String())
        }
    }
    
    return ifaces, nil
}

// 获取系统可用的的端口号， 如果传入的端口号可用，那就直接返回
func GetAvailablePort(port string) (int, error) {
    address, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:0", "0.0.0.0"))
    if err != nil {
        return 0, err
    }
    
    // 如果没有被占用 就直接返回
    if !checkPortIsOpen(port) {
        return strconv.Atoi(port)
    }
    
    return AvailablePort(address)
}

func AvailablePort(address *net.TCPAddr) (int, error) {
    listener, err := net.ListenTCP("tcp", address)
    if err != nil {
        return 0, err
    }
    
    defer listener.Close()
    return listener.Addr().(*net.TCPAddr).Port, nil
}

func checkPortIsOpen(port string) bool {
    timeout := time.Second
    conn, err := net.DialTimeout("tcp", net.JoinHostPort("127.0.0.1", port), timeout)
    if err != nil {
        return false
    }
    if conn != nil {
        defer conn.Close()
        return true
    }
    return false
}

func GetLocalIP() (ip string, err error) {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return
    }
    for _, addr := range addrs {
        ipAddr, ok := addr.(*net.IPNet)
        if !ok {
            continue
        }
        if ipAddr.IP.IsLoopback() {
            continue
        }
        if !ipAddr.IP.IsGlobalUnicast() {
            continue
        }
        return ipAddr.IP.String(), nil
    }
    return
}

func GetOutBoundIP() (ip string, err error) {
    conn, err := net.Dial("udp", "8.8.8.8:53")
    if err != nil {
        return
    }
    localAddr := conn.LocalAddr().(*net.UDPAddr)
    ip = strings.Split(localAddr.String(), ":")[0]
    return
}
