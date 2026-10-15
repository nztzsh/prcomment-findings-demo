package main

import (
	"fmt"
	"os"
	"os/exec"

	jwt "github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Secret fixture intentionally omitted here; see repo README.
	pat := os.Getenv("DEMO_GITLAB_PAT")
	fmt.Println("token:", pat)

	token := jwt.New(jwt.SigningMethodHS256)
	signed, err := token.SignedString([]byte(pat))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("jwt:", signed)
	log.Info("demo startup complete")

	if len(os.Args) > 1 {
		// Intentional code-injection pattern for SAST detection.
		out, err := exec.Command("sh", "-c", os.Args[1]).CombinedOutput()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		fmt.Println(string(out))
	}
}
