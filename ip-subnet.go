package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func main() {
	ipStr := "192.168.1.1/24"
	ip, subnet, err := net.ParseCIDR(ipStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("IP Address:", ip)
	fmt.Println("Subnet Mask:", maskToDecimal(subnet.Mask))
}

func maskToDecimal(mask net.IPMask) string {
	var result []string
	for _, b := range mask {
		result = append(result, strconv.Itoa(int(b)))
	}
	return strings.Join(result, ".")
}
