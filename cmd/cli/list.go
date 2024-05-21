package main

import (
	"fmt"
	"os/user"

	"github.com/SadS4ndWiCh/apm/internal/consola"
	"github.com/SadS4ndWiCh/apm/internal/goapch"
)

func ListCommand() int {
	u, err := user.Current()
	if err != nil || u.Username != "root" {
		fmt.Println("[-] you must be root to use this command")
		return 1
	}

	projects, err := goapch.GetAllProjects()
	if err != nil {
		fmt.Println("[-] failed to find all projects")
		return 1
	}

	table := consola.NewTableDefault()
	table.InsertRow([]string{"Project", "Status"})

	for _, project := range projects {
		table.InsertRow(project)
	}

	fmt.Println(table)
	return 0
}
