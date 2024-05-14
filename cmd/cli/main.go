package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(`
Welcome to the Apache Project Manager (APM)

Usage:

apm create <project name> - create a new project
apm delete <project name> - delete an existing project
apm start  <project name> - start an existing project
apm stop   <project name> - stop an existing project of running
apm restart               - restart the currently running project
apm list                  - list all available projects
	    `)

		os.Exit(0)
	}

	switch os.Args[1] {
	case "create":
		os.Exit(CreateCommand())
	case "delete":
		os.Exit(DeleteCommand())
	case "list":
		os.Exit(ListCommand())
	case "start":
		os.Exit(StartCommand())
	case "stop":
		os.Exit(StopCommand())
	case "restart":
		os.Exit(RestartCommand())
	}
}
