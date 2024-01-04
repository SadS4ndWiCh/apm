package list

import (
  "os"
  "fmt"
  "slices"
  "strings"

  "github.com/SadS4ndWiCh/apm/config"
  "github.com/SadS4ndWiCh/apm/lib/apache"
)

var excludeProjects []string = []string { "000-default.conf", "default-ssl.conf" }

func ListProjectsCommand() int {
  files, err := os.ReadDir(config.APACHE_SITES_ROOT)
  if err != nil {
    fmt.Println("Failed to find all projects")
    return 1
  }

  runningProject, _ := apache.CurrentRunning()

  fmt.Println("Available Projects:\n")
  for _, file := range files  {
    filename := file.Name()
    if !strings.HasSuffix(filename, ".conf") || slices.Contains(excludeProjects, filename) { continue }

    filename = strings.Replace(filename, ".conf", "", 1)
    if runningProject != "" && filename == runningProject {
      filename = fmt.Sprintf("%s - [RUNNING]", filename)
    }

    fmt.Println(filename)
  }

  return 0
}
