package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	sourcePath := "/chaosd"
	destPath := "/var/lib/chaosd/chaosd"

	err := os.MkdirAll("/var/lib/chaosd", 0755)
	if err != nil {
		fmt.Printf("Error while creating /var/lib/chaosd: %v\n", err)
		os.Exit(1)
	}

	input, err := os.ReadFile(sourcePath)
	if err != nil {
		fmt.Printf("Error while reading %s: %v\n", sourcePath, err)
		os.Exit(1)
	}

	err = os.WriteFile(destPath, input, 0755)
	if err != nil {
		fmt.Printf("Error while writing %s: %v\n", destPath, err)
		os.Exit(1)
	}

	err = os.MkdirAll("/etc/chaosd/certs", 0755)
	if err != nil {
		fmt.Printf("Error while creating /etc/chaosd/certs: %v\n", err)
		os.Exit(1)
	}

	// Lancer chaosd server
	cmd := exec.Command(destPath, "server",
		"--cert", "/etc/chaosd/certs/chaosd.crt",
		"--key", "/etc/chaosd/certs/chaosd.key")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error while executing chaosd: %v\n", err)
		os.Exit(1)
	}
}
