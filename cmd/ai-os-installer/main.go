package main

import (
	"fmt"
	"os"
)

const version = "0.1.0-dev"

func main() {
	fmt.Printf("AI OS Installer v%s\n", version)
	fmt.Println("⚠️  MVP in development - Not ready for use yet")
	os.Exit(0)
}
