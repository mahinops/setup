package cmd

import (
	"fmt"
	"os"
	"setup/internal/system"
)

// ANSI escape codes for styling
const (
	red   = "\033[31;1m" // Bold Red
	reset = "\033[0m"    // Reset color
)

func Execute() {
	if len(os.Args) < 2 {
		showHelp()
		return
	}

	arg := os.Args[1]
	switch arg {
	case "-h", "--help":
		showHelp()
	case "-s", "--sysinfo":
		system.ShowSystemInfo()
	case "-l", "--list":
		system.ListSoftwares()
	case "-u", "--update":
		system.UpdatePackages()
	default:
		// Bigger, bolder alert
		fmt.Println(red + "🚨 ERROR: Invalid command! 🚨" + reset)
		fmt.Println(red + "❌ Command not recognized: " + arg + reset)
		showHelp()
	}
}

func showHelp() {
	fmt.Println("Usage:")
	fmt.Println("  -s, --sysinfo   Show system information")
	fmt.Println("  -l, --list      List installed software")
	fmt.Println("  -u, --update    Run 'sudo apt-get update'")
	fmt.Println("  -h, --help      Show available commands")
	os.Exit(1)
}
