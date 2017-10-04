package duplicacy

import (
	"fmt"
	"os"
	"os/exec"
)

func Revisions() {
	var (
		cmdOut []byte
		err    error
	)
	cmdName := "duplicacy"
	cmdArgs := []string{"list"}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running duplicacy : ", err)
		os.Exit(1)
	}
	output := string(cmdOut)
	fmt.Println(output)
}
