package main

import (
	"fmt"
	"os"

	"github.com/Abu-Zakaria/process-tracker/pkg/capture_processes"
)

const (
	memory_data_file_path = "./mem_data.json"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Starting process tracker...")
		capture_processes.Run(memory_data_file_path)
	}
}
