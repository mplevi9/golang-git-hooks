package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	files := os.Args[1:]

	for _, f := range files {
		ext := f[len(f)-3:]
		if ext != ".go" {
			continue
		}

		cmd := exec.Command("golangci-lint", "run", f)
		out, err := cmd.CombinedOutput()

		if err != nil {
			log.Fatal(fmt.Sprintf("%s\n", out))
		}

		fmt.Printf("%s\n", out)
	}

}
