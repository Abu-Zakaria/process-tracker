package main

import (
	"fmt"
	"os"

	"github.com/Abu-Zakaria/process-tracker/pkg/capture_processes"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Starting process tracker...")
		capture_processes.Run()
	}
}
