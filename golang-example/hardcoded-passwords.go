package main

import "fmt"

func main() {
	// 🚨 Vulnerability: Hardcoded password
	password := "SuperSecret123"
	fmt.Println("Your password is:", password)
}
