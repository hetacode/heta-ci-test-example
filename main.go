package main

import (
	"fmt"
	"os"
)

func main() {
	hostanme, _ := os.Hostname()
	fmt.Println("develop")
	fmt.Println("hostname: " + hostanme)
	fmt.Println("from develop branch")
}
