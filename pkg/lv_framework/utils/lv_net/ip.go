package lv_net

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
)

// GET preferred outbound ip of this machine
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8")
	if err != nil {
		return ""
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String())
	return localAddr.IP.String()
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

// IsPrivateIP 判断是否为内网IP
func IsPrivateIP(ipStr string) bool {
	// 将字符串解析为 net.IP
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return false // 如果解析失败，返回 false
	}

	// 定义内网IP地址段
	privateIPBlocks := []*net.IPNet{
		{IP: net.ParseIP("10.0.0.0"), Mask: net.CIDRMask(8, 32)},
		{IP: net.ParseIP("172.16.0.0"), Mask: net.CIDRMask(12, 32)},
		{IP: net.ParseIP("192.168.0.0"), Mask: net.CIDRMask(16, 32)},
		{IP: net.ParseIP("127.0.0.0"), Mask: net.CIDRMask(8, 32)},
		{IP: net.ParseIP("169.254.0.0"), Mask: net.CIDRMask(16, 32)},
	}

	// 遍历所有内网地址段，检查IP是否在其中
	for _, block := range privateIPBlocks {
		if block.Contains(ip) {
			return true
		}
	}

	return false
}

func GetClientRealIP(c *gin.Context) string {
	// 获取用户请求的 IP 地址
	ip := c.Request.Header.Get("X-Real-IP")
	if ip == "" {
		ip = c.Request.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip = c.Request.RemoteAddr
	}
	return ip
}

// 获取外网ip地址
func GetLocation(ip string) string {
	return ip
	//if ip == "127.0.0.1" || ip == "localhost" {
	//	return "内部IP"
	//}
	//resp, err := http.Get("https://restapi.amap.com/v3/ip?ip=" + ip + "&key=3fabc36c20379fbb9300c79b19d5d05e")
	//if err != nil {
	//	panic(err)
	//
	//}
	//defer resp.Body.Close()
	//s, err := ioutil.ReadAll(resp.Body)
	//fmt.Printf(string(s))
	//
	//m := make(map[string]string)
	//
	//err = json.Unmarshal(s, &m)
	//if err != nil {
	//	fmt.Println("Umarshal failed:", err)
	//}
	//if m["province"] == "" {
	//	return "未知位置"
	//}
	//return m["province"] + "-" + m["city"]
}

// 获取局域网ip地址
func GetLocaHost() string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}

	}
	return ""
}
