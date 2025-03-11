package system

import (
	"fmt"
	"runtime"
)

// ShowSystemInfo prints OS and architecture details
func ShowSystemInfo() {
	fmt.Println("System Information:")
	fmt.Printf("OS: %s\n", runtime.GOOS)
	fmt.Printf("Arch: %s\n", runtime.GOARCH)
}
