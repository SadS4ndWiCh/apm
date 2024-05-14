package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/SadS4ndWiCh/apm/internal/goapch"
)

func ListCommand() int {
	u, err := user.Current()
	if err != nil || u.Username != "root" {
		fmt.Println("[-] you must be root to use this command")
		return 1
	}

	files, err := os.ReadDir(goapch.APACHE_SITES_ROOT)
	if err != nil {
		fmt.Println("[-] failed to find all projects")
		return 1
	}

	runningProject, _ := goapch.CurrentRunning()
	fmt.Println(runningProject)

	fmt.Println("+-----------------+---------+")
	fmt.Printf("| %-15s | %-7s |\n", "Project", "Status")
	fmt.Println("+-----------------+---------+")
	for _, file := range files {
		filename := strings.Replace(file.Name(), ".conf", "", 1)
		if filename == "000-default" || filename == "default-ssl" {
			continue
		}

		status := ""
		if filename == runningProject {
			status = "RUNNING"
		}

		fmt.Printf("| %-15s | %-7s |\n", filename, status)
		fmt.Println("+-----------------+---------+")
	}

	return 0
}
