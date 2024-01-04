package start

import (
  "os"
  "fmt"

  "github.com/SadS4ndWiCh/apm/lib/apache"
)

func StartProjectCommand() int {
  if len(os.Args) == 2 {
    fmt.Println("'start' command must require the project name")
    return 1
  }

  if running, err := apache.CurrentRunning(); err == nil {
    fmt.Printf("Stopping currently running '%s' project", running)
    apache.Stop(running)
  }

  projectName := os.Args[2]

  if _, err := apache.Start(projectName); err != nil {
    fmt.Printf("Failed to start the '%s' project: %v", projectName, err)
    return 1
  }

  if err := apache.Restart(); err != nil {
    fmt.Println("Failed to restart the apache")
    return 1
  }

  fmt.Printf("Project '%s' started successfuly", projectName)

  return 0
}
