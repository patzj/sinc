package main

import (
	"fmt"

	"github.com/patzj/network-calculator/models"
)

func printIpV4SubnetHeader() {
	fmt.Println("---------------------------------")
	fmt.Println("|          IPv4 Subnet          |")
	fmt.Println("---------------------------------")
}

func promptIpV4SubnetInput() {
	ipv4Str := ""
	fmt.Print("Enter IPv4 address (e.g. 192.168.0.1): ")
	fmt.Scanln(&ipv4Str)

	ipv4, err := models.NewIPv4(ipv4Str)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(ipv4)

	cidrStr := ""
	fmt.Print("Enter subnet mask (0-32): ")
	fmt.Scanln(&cidrStr)

	netmask, err := models.NewNetmask(cidrStr)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(netmask)
}
