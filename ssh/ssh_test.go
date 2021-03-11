package ssh

import (
	"os"
	"testing"
)

func TestSsh(t *testing.T) {
	client := Client{
		Config: Config{
			User:         "ubuntu",
			Password:     "yufield101",
			Host:         "192.168.214.205",
			Port:         "8822",
			IdentityFile: "",
		},
		Stdout:         os.Stdout,
		Stderr:         os.Stdout,
		ConnectRetries: 0,
		client:         nil,
	}
	err := client.Connect()
	if err != nil {
		t.Error(err)
	}
	_ = client.Exec("git clone https://github.com/fdev-ci/mock-demo.git")
}
