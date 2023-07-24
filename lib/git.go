package lib

import (
	"fmt"
	"os/exec"
	"strings"
)

func executeCommand(command string, path string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Dir = path

	output, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
	}

	return string(output), err
}

func GetGitProjectName(path string) (string, error) {

	return executeCommand("git config --local remote.origin.url|sed -n 's#.*/\\([^.]*\\)\\.git#\\1#p'", path)
}

func GetGitCommitData(path string, dateStart string, dateUntil string) ([]string, error) {
	var result string
	if dateStart == "" {
		result, _ = executeCommand("git log --format=%B -n 1", path)
	} else if dateUntil == "" {
		result, _ = executeCommand("git log --format=%B --since='"+dateStart+" 00:00'", path)
	} else {
		result, _ = executeCommand("git log --format=%B --since='"+dateStart+" 00:00' --until='"+dateUntil+" 23:59'", path)
	}

	return strings.Split(result, "\n\n"), nil
}
