package restart

import (
  "fmt"

  "github.com/SadS4ndWiCh/apm/lib/apache"
)

func RestartApacheCommand() int {
  if err := apache.Restart(); err != nil {
    fmt.Println("Failed to restart the apache")
    return 1
  }

  fmt.Println("Apache was successfuly restarted")
  return 0
}
