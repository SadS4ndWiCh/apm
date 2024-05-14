package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/SadS4ndWiCh/apm/internal/goapch"
)

func StopCommand() int {
	u, err := user.Current()
	if err != nil || u.Username != "root" {
		fmt.Println("[-] you must be root to use this command")
		return 1
	}

	if len(os.Args) < 3 {
		fmt.Println("[-] missing <project name> argument")
		return 1
	}

	projectName := os.Args[2]
	if _, err := goapch.Stop(projectName); err != nil {
		fmt.Printf("[-] failed to stop the `%s` project\n", projectName)
		return 1
	}

	fmt.Printf("[+] stopped '%s' project\n", projectName)
	return 0
}
