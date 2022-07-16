package goip

import (
	"fmt"
	"net"
)

func Get() {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, netInterface := range netInterfaces {
		addrsList, err := netInterface.Addrs()
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, addrs := range addrsList {
			if ipnet, ok := addrs.(*net.IPNet); ok {
				if ipnet.IP.To4() != nil {
					fmt.Println(ipnet.IP.String())
				}
			}
		}
	}
}
