package sysutil

import (
	"encoding/base64"
	"fmt"
	"net"
)

func useDemo() {
	mac := GetMac()
	a := "MWU6MDA6ZGE6MzQ6NWE6NTI="
	decoded, _ := base64.StdEncoding.DecodeString(a)
	decodeStr := string(decoded)
	if mac != decodeStr {
		fmt.Println("fail")
	} else {
		fmt.Println("success")
	}
}

func GetMac() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, inter := range interfaces {
		mac := inter.HardwareAddr
		if mac.String() != "" {
			return mac.String()
		}
	}
	return ""
}
