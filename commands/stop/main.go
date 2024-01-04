package stop

import (
  "fmt"

  "github.com/SadS4ndWiCh/apm/lib/apache"
)

func StopProjectCommand() int {
  projectName, err := apache.CurrentRunning()
  if err != nil {
    fmt.Println("No one project currently running")
    return 1
  }

  if _, err := apache.Stop(projectName); err != nil {
    fmt.Printf("Failed to stop the '%s' project: %v", projectName, err)
    return 1
  }

  if err := apache.Restart(); err != nil {
    fmt.Println("Failed to restart the apache")
    return 1
  }

  fmt.Printf("Project '%s' stoped successfuly", projectName)

  return 0
}
