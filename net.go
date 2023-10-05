package goutils

import (
	"net"
	"time"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/5 18:48
  @describe :
*/

// GetMasterIP 获取连接互联网的主IP地址
func GetMasterIP() string {
	conn, err := net.DialTimeout("udp", "8.8.8.8:8", time.Second*5)
	if err != nil {
		return ""
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

// GetLocalIPv4 获取当前IPv4地址
func GetLocalIPv4() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && !ipNet.IP.IsLinkLocalUnicast() && ipNet.IP.To4() != nil {
			return ipNet.IP.To4().String()
		}
	}
	return ""
}

// GetLocalIPv6 获取当前IPv6地址
func GetLocalIPv6() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && !ipNet.IP.IsLinkLocalUnicast() && ipNet.IP.To16() != nil {
			return ipNet.IP.To16().String()
		}
	}
	return ""
}
