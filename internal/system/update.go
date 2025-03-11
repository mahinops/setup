package system

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// UpdatePackages runs 'sudo apt-get update' and shows real-time logs
func UpdatePackages() {
	fmt.Println("\nUpdating package lists...")
	cmd := exec.Command("sudo", "apt-get", "update")

	// Redirect real-time logs to terminal
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error executing update command: %v\n", err)
	}

	fmt.Println("Update completed successfully.")
}
