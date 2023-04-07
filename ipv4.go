package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/patzj/sinc/models"
	"github.com/patzj/sinc/utils"
)

func printIpV4SubnetHeader() {
	fmt.Println("---------------------------------")
	fmt.Println("|          IPv4 Subnet          |")
	fmt.Println("---------------------------------")
}

func promptIpV4SubnetInput() {
	ipv4 := new(models.IPv4)
	promptAddressInput(ipv4)

	netmask := new(models.Netmask)
	promptNetmaskInput(netmask)

	if networkAddr, broadcastAddr, err := ipv4.Subnet(*netmask); err != nil {
		fmt.Println(err.Error())
	} else {
		subnetResult(networkAddr, broadcastAddr)
	}
}

func promptAddressInput(ipv4 *models.IPv4) {
	for {
		ipv4Str := ""
		fmt.Print("Enter IPv4 address (e.g. 192.168.0.1): ")
		fmt.Scanln(&ipv4Str)

		if result, err := models.NewIPv4(ipv4Str); err != nil {
			fmt.Println(err.Error())
		} else {
			*ipv4 = *result
			break
		}
	}
}

func promptNetmaskInput(netmask *models.Netmask) {
	for {
		cidrStr := ""
		fmt.Print("Enter subnet mask (0-32): ")
		fmt.Scanln(&cidrStr)

		if result, err := models.NewNetmask(cidrStr); err != nil {
			fmt.Println(err.Error())
		} else {
			*netmask = *result
			break
		}
	}
}

func subnetResult(start, end *models.IPv4) {
	fmt.Printf("Subnet range: %s - %s\n", start.Octets(), end.Octets())

	for {
		fmt.Println("0 - Main menu")
		fmt.Println("1 - Export")

		choice := 0
		promptChoice(&choice)

		switch choice {
		case 0:
			goto Exit
		case 1:
			exportSubnetToJson(start, end)
			goto Exit
		default:
			fmt.Println("Invalid choice")
		}

		fmt.Println()
	}

Exit:
}

func exportSubnetToJson(start, end *models.IPv4) {
	hosts := utils.Ipv4Hosts(start, end)
	hostsStr := []string{}

	for _, host := range hosts {
		hostsStr = append(hostsStr, host.Octets().String())
	}

	result := models.IPv4SubnetResult{
		Network:   start.Octets().String(),
		Broadcast: end.Octets().String(),
		Hosts:     hostsStr,
	}

	data, _ := json.MarshalIndent(result, "", "  ")
	filepath := path.Join(os.TempDir(), "ipv4_subnet.json")
	_ = ioutil.WriteFile(filepath, data, 0644)
}
