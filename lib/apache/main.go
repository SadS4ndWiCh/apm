package apache

import (
  "os/exec"
  "strings"
)

func Start(project string) (out []byte, err error) {
  startCmd := exec.Command("a2ensite", project, "-q")
  out, err = startCmd.Output()

  return
}

func Stop(project string) (out []byte, err error) {
  stopCmd := exec.Command("a2dissite", project, "-q")
  out, err = stopCmd.Output()

  return
}

func CurrentRunning() (string, error) {
  stopCmd := exec.Command("a2query", "-s")
  out, err := stopCmd.Output()

  if err != nil { return "", err }

  project := strings.Fields(string(out))[0]

  return project, nil
}

func Restart() error {
  restartCmd := exec.Command("service", "apache2", "restart")
  _, err := restartCmd.Output()

  return err
}
