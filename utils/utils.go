package utils

import (
	"os"
	"os/exec"
)

func ClearTerminal() {
	cmd := exec.Command("clear")
	if os.PathSeparator == '\\' {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}