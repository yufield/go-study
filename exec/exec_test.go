package exec

import (
	"os"
	"os/exec"
	"testing"
)

func TestCmd(t *testing.T) {
	cmd := exec.Command("cmd", "/c", "echo %my_var%")
	cmd.Stdout = os.Stdout
	cmd.Env = append(os.Environ(), "MY_VAR=some_value")
	err := cmd.Run()
	if err != nil {
		t.Error(err)
	}
}
