package utilities

import (
	"os/exec"
)

func CommandLine(ps string) ([]byte, error) {
	cmd := exec.Command("powershell", "-NoProfile", "-NonInteractive", "-Command", ps)
	return cmd.Output()
}
