package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Added in pr-redesigned: a second intentional injection pattern in the PR diff.
func runUserCommandRedesigned() {
	if len(os.Args) > 2 {
		out, _ := exec.Command("bash", "-c", os.Args[2]).CombinedOutput()
		fmt.Println(string(out))
	}
}
