package main

import (
  "fmt"
  "os"

  "github.com/SadS4ndWiCh/apm/commands/restart"
  "github.com/SadS4ndWiCh/apm/commands/create"
  "github.com/SadS4ndWiCh/apm/commands/delete"
  "github.com/SadS4ndWiCh/apm/commands/start"
  "github.com/SadS4ndWiCh/apm/commands/stop"
  "github.com/SadS4ndWiCh/apm/commands/list"
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
    os.Exit(create.CreateProjectCommand())
  case "delete":
    os.Exit(delete.DeleteProjectCommand())
  case "list":
    os.Exit(list.ListProjectsCommand())
  case "start":
    os.Exit(start.StartProjectCommand())
  case "stop":
    os.Exit(stop.StopProjectCommand())
  case "restart":
    os.Exit(restart.RestartApacheCommand())
  default:
    fmt.Printf("Unknow command: %s", os.Args[1])
    os.Exit(1)
  }
}
