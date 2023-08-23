package main

import (
	"os/exec"
)

func main() {
	cmd1 := exec.Command("calc;explorer -")

	cmd1.Run()
}
