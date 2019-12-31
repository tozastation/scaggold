package tree

import (
	"os"
	"os/exec"
)

func Exec() error {
	treeCmd := exec.Command("tree") // ...enable us to pass them slice
	treeCmd.Stdout = os.Stdout
	treeCmd.Stderr = os.Stderr
	if err := treeCmd.Run(); err != nil {
		return err
	}
	return nil
}
