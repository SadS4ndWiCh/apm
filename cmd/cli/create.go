package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/SadS4ndWiCh/apm/internal/goapch"
)

func CreateCommand() int {
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

	alreadyExists, err := goapch.IsProjectExists(projectName)
	if err != nil {
		fmt.Println("[-] failed to find all projects")
		return 1
	}

	if alreadyExists {
		fmt.Println("[-] project already exists")
		return 1
	}

	projectFolderPath := fmt.Sprintf("%s/%s", goapch.APACHE_PROJECT_ROOT, projectName)
	projectConfigFile := fmt.Sprintf(`
<VirtualHost *:80>
    ServerName %s
    ServerAlias %s

    DocumentRoot %s
    ErrorLog ${APACHE_LOG_DIR}/error.log
    CustomLog ${APACHE_LOG_DIR}/access.log combined
</VirtualHost>
  `, projectName, projectName, projectFolderPath)
	projectConfigPath := fmt.Sprintf("%s/%s.conf", goapch.APACHE_SITES_ROOT, projectName)

	if err := os.WriteFile(projectConfigPath, []byte(projectConfigFile), 0644); err != nil {
		fmt.Printf("[-] failed to create '%s' config file\n", projectConfigPath)
		fmt.Println(err)
		return 1
	}

	if err := os.Mkdir(projectFolderPath, 0644); err != nil {
		fmt.Printf("[-] failed to create '%s' project folder\n", projectFolderPath)
		return 1
	}

	// Change the owner from project folder path to regular user
	if err := os.Chown(projectFolderPath, 1000, 1000); err != nil {
		fmt.Println("[*] failed to change the owner of project folder")
		fmt.Printf("[*] change manually with `sudo chown $USER:$USER %s`\n", projectFolderPath)
	}

	fmt.Printf("[+] Project '%s' was created\n", projectName)
	return 0
}
