package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/SadS4ndWiCh/apm/internal/goapch"
)

func DeleteCommand() int {
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
	projectFolderPath := fmt.Sprintf("%s/%s", goapch.APACHE_PROJECT_ROOT, projectName)
	projectConfigPath := fmt.Sprintf("%s/%s.conf", goapch.APACHE_SITES_ROOT, projectName)

	if err := os.RemoveAll(projectFolderPath); err != nil {
		fmt.Println("[*] failed to remove project folder")
		fmt.Printf("[*] manually delete with `rm -fr %s`\n", projectFolderPath)
	}

	if err := os.Remove(projectConfigPath); err != nil {
		fmt.Println("[*] failed to remove config file")
		fmt.Printf("[*] manually delete with `sudo rm -fr %s`\n", projectConfigPath)
	}

	fmt.Printf("[+] Project '%s' was deleted\n", projectName)
	return 0
}
