// Copyright (c) 2016, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package membergroup

import (
	"fmt"
	"net"
	"strings"
)

// getLocalIP .
func getLocalIP() (ip string) {
	interfaces, err := net.Interfaces()
	if nil != err {
		return ""
	}

	for _, iface := range interfaces {
		if !interfaceNameFilter(iface.Name) {
			continue
		}

		addrs, err := iface.Addrs()
		if nil != err {
			fmt.Println(err.Error())
			continue
		}

		for _, address := range addrs {
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if nil != ipnet.IP.To4() {
					ip = ipnet.IP.String()
					//break
				}
			}
		}
	}

	return ip
}

// interfaceNameFilter .
func interfaceNameFilter(name string) bool {
	return strings.HasPrefix(name, "eth") || strings.HasPrefix(name, "ens")
}
