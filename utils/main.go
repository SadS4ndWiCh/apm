package utils

import (
  "os"
  "strings"

  "github.com/SadS4ndWiCh/apm/config"
)

func IsProjectExists(projectName string) (bool, error) {
  files, err := os.ReadDir(config.APACHE_SITES_ROOT)
  if err != nil { return false, err }

  for _, file := range files {
    if filename := file.Name(); strings.HasSuffix(filename, ".conf") && strings.Replace(filename, ".conf", "", 1) == projectName {
      return true, nil
    }
  }

  return false, nil
}
