package main

import "fmt"

func printIpV4SubnetHeader() {
	fmt.Println("---------------------------------")
	fmt.Println("|          IPv4 Subnet          |")
	fmt.Println("---------------------------------")
}

func promptIpV4SubnetInput() {
	ipv4 := ""
	fmt.Print("Enter IPv4 address (e.g. 192.168.0.1): ")
	fmt.Scanln(&ipv4)

	subnet := ""
	fmt.Print("Enter subnet mask (1-32): ")
	fmt.Scanln(&subnet)
}
