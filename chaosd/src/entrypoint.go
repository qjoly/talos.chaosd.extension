package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	sourcePath := "/chaosd"
	destPath := "/var/lib/chaosd/chaosd"

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

	var cmd *exec.Cmd
	if os.Getenv("CHAOSD_CERT") != "" && os.Getenv("CHAOSD_KEY") != "" {
		cmd = exec.Command(destPath, "server",
			"--cert", os.Getenv("CHAOSD_CERT"),
			"--key", os.Getenv("CHAOSD_KEY"))
	} else {
		cmd = exec.Command(destPath, "server")
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error while executing chaosd: %v\n", err)
		os.Exit(1)
	}
}
