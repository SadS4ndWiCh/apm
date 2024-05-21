package goapch

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	APACHE_SITES_ROOT         = "/etc/apache2/sites-available"
	APACHE_SITES_DEFAULT_CONF = "/etc/apache2/sites-available/000-default.conf"
	APACHE_PROJECT_ROOT       = "/var/www"
)

func Start(project string) (out []byte, err error) {
	out, err = exec.Command("a2ensite", project, "-q").Output()
	return
}

func Stop(project string) (out []byte, err error) {
	out, err = exec.Command("a2dissite", project, "-q").Output()
	return
}

func Restart() (err error) {
	_, err = exec.Command("service", "apache2", "restart").Output()
	return
}

func CurrentRunning() (string, error) {
	out, err := exec.Command("a2query", "-s").Output()
	if err != nil {
		return "", err
	}

	fields := strings.Fields(string(out))
	if len(fields) == 0 {
		return "", errors.New("not found current running project")
	}

	return fields[0], nil
}

func GetAllProjects() ([][]string, error) {
	files, err := os.ReadDir(APACHE_SITES_ROOT)
	if err != nil {
		return nil, errors.New("failed to find all projects")
	}

	projects := [][]string{}
	runningProject, _ := CurrentRunning()

	for _, file := range files {
		filename := strings.Replace(file.Name(), ".conf", "", 1)
		if filename == "000-default" || filename == "default-ssl" {
			continue
		}

		status := "-"
		if filename == runningProject {
			status = "RUNNING"
		}

		projects = append(projects, []string{filename, status})
	}

	return projects, nil
}

func IsProjectExists(projectName string) (bool, error) {
	files, err := os.ReadDir(APACHE_SITES_ROOT)
	if err != nil {
		return false, err
	}

	projectName = fmt.Sprintf("%s.conf", projectName)

	for _, file := range files {
		if file.Name() == projectName {
			return true, nil
		}
	}

	return false, nil
}
