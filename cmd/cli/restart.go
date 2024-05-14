package main

import (
	"fmt"
	"os/user"

	"github.com/SadS4ndWiCh/apm/internal/goapch"
)

func RestartCommand() int {
	u, err := user.Current()
	if err != nil || u.Username != "root" {
		fmt.Println("[-] you must be root to use this command")
		return 1
	}

	if err := goapch.Restart(); err != nil {
		fmt.Println("[-] failed to restart apache")
		return 1
	}

	fmt.Println("[+] apache restarted")
	return 0
}
