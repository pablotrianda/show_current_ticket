package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const COMMAND = "ps -aux | grep mattermost | wc -l"
const GIT_GET_NAME = "git rev-parse --abbrev-ref HEAD"

func main() {
	// Expect a git path to get the current branch
	proyect_path := os.Args[1]

	showMessage := ""
	qt_lines := runCommand(COMMAND)
	process, _ := strconv.Atoi(qt_lines)

	// If mattermost isn't running the message is empty
	if process <= 2 {
		fmt.Println(showMessage)
		return
	}

	// Show the name of the current branch(TICKET-NUMBER)
	os.Chdir(proyect_path)
	branchName := runCommand(GIT_GET_NAME)

	fmt.Println("🔥 👨‍💻 " + branchName)

}

// Run a bash command
func runCommand(command string) string {
	output, _ := exec.Command("bash", "-c", command).CombinedOutput()
	return cleanString(string(output))
}

func cleanString(s string) string {
	return strings.ReplaceAll(string(s), "\n", "")
}
