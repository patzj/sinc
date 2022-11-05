package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		printHeader()
		printOptions()

		choice := 0
		promptChoice(&choice)

		switch choice {
		case 0:
			goto Exit
		case 1:
			printIpV4SubnetHeader()
			promptIpV4SubnetInput()
		case 2:
			printIpV6PrefixHeader()
			promptIpV6PrefixInput()
		default:
			fmt.Println("Invalid choice")
		}

		fmt.Println()
	}

Exit:
	fmt.Println("Goodbye!")
	os.Exit(0)
}

func printHeader() {
	fmt.Println("----------------------------------------")
	fmt.Println("|          Network Calculator          |")
	fmt.Println("----------------------------------------")
}

func printOptions() {
	fmt.Println("0 - Exit")
	fmt.Println("1 - IPv4 subnet")
	fmt.Println("2 - IPv6 prefix")
}

func promptChoice(choice *int) {
	fmt.Print("Enter choice: ")
	if _, err := fmt.Scanf("%d", choice); err != nil {
		os.Exit(1)
	}
	fmt.Println()
}
