package main

import (
	"fmt"
	"os"
)

func main() {
	promptFileToRead()
}

func printHeader() {
	fmt.Println("-------------------------------------------------")
	fmt.Println("|           Simple Network Calculator           |")
	fmt.Println("-------------------------------------------------")
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

func promptFileToRead() {
	fmt.Print("Enter choice: ")
	var choice string
	if _, err := fmt.Scanf("%s", &choice); err != nil {
		os.Exit(1)
	}
	fileRead(choice)
	fmt.Println()
}

func fileRead(filepath string) {
	dat, err := os.ReadFile(filepath)
	if err != nil {
		os.Exit(1)
	} else {
		fmt.Println(string(dat))
	}
}
