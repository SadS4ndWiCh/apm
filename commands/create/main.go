package create

import (
  "os"
  "os/user"
  "fmt"
  "strconv"

  "github.com/SadS4ndWiCh/apm/config"
  "github.com/SadS4ndWiCh/apm/utils"
)

func CreateProjectCommand() int {
  if len(os.Args) == 2 {
    fmt.Println("'create' command must require the project name")

    return 1
  }

  u, err := user.Current()
  if err != nil || u.Username != "root" {
    fmt.Println("[ERROR] Must be root user to use this command")
    return 1
  }

  projectName := os.Args[2]
  if exists, err := utils.IsProjectExists(projectName); exists || err != nil {
    fmt.Println("[ERROR] Project already exists")
    return 1
  }

  projectFolderPath := fmt.Sprintf("%s/%s", config.APACHE_PROJECT_ROOT, projectName)
  projectConfigFile := fmt.Sprintf(`
<VirtualHost *:80>
    ServerName %s
    ServerAlias %s

    DocumentRoot %s
    ErrorLog \${APACHE_LOG_DIR}/error.log
    CustomLog \${APACHE_LOG_DIR}/access.log combined
</VirtualHost>
  `, projectName, projectName, projectFolderPath)
  projectConfigPath := fmt.Sprintf("%s/%s.conf", config.APACHE_SITES_ROOT, projectName)

  if err := os.WriteFile(projectConfigPath, []byte(projectConfigFile), os.ModePerm); err != nil {
    fmt.Printf("[ERROR] Failed to create config file: %v\n", err)
    return 1
  }

  var uid, gid int
  var _err error

  uid, _err = strconv.Atoi(os.Getenv("SUDO_UID"))
  if _err != nil {
    fmt.Printf("[ERROR] Failed to convert uid: %v\n", _err)
    return 1
  }

  gid, _err = strconv.Atoi(os.Getenv("SUDO_GID"))
  if _err != nil {
    fmt.Printf("[ERROR] Failed to convert gid: %v\n", _err)
    return 1
  }

  if err := os.Chown(projectConfigPath, uid, gid); err != nil {
    fmt.Printf("[ERROR] Failed to change config file ownership: %v\n", err)
    return 1
  }

  if err := os.Mkdir(projectFolderPath, os.ModePerm); err != nil {
    fmt.Printf("[ERROR] Failed project folder: %v\n", err)
    return 1
  }

  fmt.Printf("Project '%s' created successfuly", projectName)

  return 0
}
