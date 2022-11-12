package main

import (
	"fmt"
	"regexp"
)

func printIpV4SubnetHeader() {
	fmt.Println("---------------------------------")
	fmt.Println("|          IPv4 Subnet          |")
	fmt.Println("---------------------------------")
}

func promptIpV4SubnetInput() {
	ipv4 := ""
	fmt.Print("Enter IPv4 address (e.g. 192.168.0.1): ")
	fmt.Scanln(&ipv4)

	// Syntactic validation
	ipv4Pattern := `^(\d{1,3}\.){3}\d{1,3}$`
	if matched, _ := regexp.MatchString(ipv4Pattern, ipv4); !matched {
		fmt.Println("Invalid IPv4 address")
		return
	}

	subnet := ""
	fmt.Print("Enter subnet mask (1-32): ")
	fmt.Scanln(&subnet)
}
