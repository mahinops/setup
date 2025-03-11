package system

import (
	"fmt"
	"io/ioutil"
	"log"
)

// ListSoftwares prints the installed software in the 'softwares' directory
func ListSoftwares() {
	softwaresDir := "./softwares"
	files, err := ioutil.ReadDir(softwaresDir)
	if err != nil {
		log.Fatalf("Error reading 'softwares' directory: %v\n", err)
	}

	fmt.Println("\nList of installed software:")
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("- " + file.Name())
		}
	}
}
