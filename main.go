package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("test branch")
	hostanme, _ := os.Hostname()
	fmt.Println("hostname: " + hostanme)
}
