package toolbox

import (
	"fmt"
	"os/exec"
	"strings"
)

func SSH(ipAddress, command string) (string, error) {
	var s string
	query := command
	cmd, err := exec.Command("ssh", fmt.Sprintf("root@%s", ipAddress), query).CombinedOutput()
	if err != nil {
		return s, err
	}
	return strings.TrimSpace(string(cmd)), nil

}
