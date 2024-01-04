package delete

import (
  "os"
  "fmt"
  "os/user"

  "github.com/SadS4ndWiCh/apm/utils"
  "github.com/SadS4ndWiCh/apm/config"
  "github.com/SadS4ndWiCh/apm/utils/console"
)

func DeleteProjectCommand() int {
  if len(os.Args) == 2 {
    fmt.Println("'delete' command must require project name")
    return 1
  }

  if u, err := user.Current(); u.Username != "root" || err != nil {
    fmt.Println("[ERROR] Must be root to use that command")
    return 1
  }

  projectName := os.Args[2]
  if exists, err := utils.IsProjectExists(projectName); !exists || err != nil {
    fmt.Printf("[ERROR] Project \"%s\" does not exists", projectName)
    return 1
  }

  askMessage := fmt.Sprintf("Are you sure do you want delete \"%s\" project? (y/N): ", projectName)
  if canDelete, err := console.Ask(askMessage); !canDelete || err != nil {
    return 0
  }

  projectFolderPath := fmt.Sprintf("%s/%s", config.APACHE_PROJECT_ROOT, projectName)
  projectConfigPath := fmt.Sprintf("%s/%s.conf", config.APACHE_SITES_ROOT, projectName)

  if err := os.RemoveAll(projectFolderPath); err != nil {
    fmt.Println("[ERROR] Failed to remove project folder")
    return 1
  }

  if err := os.Remove(projectConfigPath); err != nil {
    fmt.Println("[ERROR] Failed to remove project config file")
    return 1
  }

  fmt.Printf("Project '%s' was successfuly deleted", projectName)

  return 0
}
