package capture_processes

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type MemoryStatus struct {
	Process Process `json:"process"`
	Memory  uint64  `json:"memory"`
}

func GetMemory(process Process) uint64 {
	cmd := exec.Command("ps", "-p", strconv.Itoa(process.Pid), "-o", "rss=")
	var output strings.Builder
	cmd.Stdout = &output

	cmd.Run()

	var mem_str string = strings.TrimRight(strings.TrimLeft(output.String(), " "), "\n")

	memory, err := strconv.ParseUint(mem_str, 10, 64)
	if err != nil {
		log.Printf("Couldn't get memory for process %d: %v", process.Pid, err)
	}

	return memory
}
