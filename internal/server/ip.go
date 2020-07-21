package server

import (
	"fmt"
	"net"
)

// IPByInterfaceName returns IP associated with interface nam
func IPByInterfaceName(interfaceName string) (string, error) {
	netIf, err := net.InterfaceByName(interfaceName)
	if err != nil {
		return "", err
	}

	ifAddrs, err := netIf.Addrs()
	if err != nil {
		return "", err
	}

	for _, ifAddr := range ifAddrs {
		var ip *net.IP
		switch addr := ifAddr.(type) {
		case *net.IPNet:
			ip = &addr.IP
		case *net.IPAddr:
			ip = &addr.IP
		default:
			return "", err
		}

		if ip == nil || ip.IsLoopback() || ip.To4() == nil {
			continue
		}

		return ip.String(), nil
	}

	return "", fmt.Errorf("no suitable ipv4 address found for %s interface", interfaceName)
}
