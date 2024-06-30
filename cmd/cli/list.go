package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/SadS4ndWiCh/apm/internal/goapch"
	"github.com/jedib0t/go-pretty/v6/table"
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

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Project", "Status"})

	for _, project := range projects {
		t.AppendRow(table.Row{project[0], project[1]})
	}

	t.Render()
	return 0
}
