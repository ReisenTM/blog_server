package ip

import "net"

func HasLocalIPAddr(ip string) bool {
	return HasLocalIP(net.ParseIP(ip))
}

// HasLocalIP 通过ip判断内网
func HasLocalIP(ip net.IP) bool {
	if ip.IsLoopback() {
		return true
	}
	ip4 := ip.To4()
	if ip4 == nil {
		return false
	}
	return ip4[0] == 10 ||
		(ip4[0] == 172 && ip4[1] >= 16 && ip[4] <= 31) ||
		(ip4[0] == 192 && ip4[1] == 168) ||
		(ip4[0] == 169 && ip4[1] == 254)
}
