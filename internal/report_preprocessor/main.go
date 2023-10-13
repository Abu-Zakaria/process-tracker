package report_preprocessor

import (
	"fmt"

	"github.com/Abu-Zakaria/process-tracker/internal/capture_processes"
	"github.com/rodaine/table"
)

type ProcessInfo struct {
	Pid        int
	Executable string
	MemorySize uint64
	Time       string
}

func RenderCapture(capture capture_processes.Capture) {
	fmt.Println("")

	fmt.Println("Process capture time:", capture.Time)
	fmt.Println("")

	PrintMemStatuses(capture.MemoryStatuses)

	fmt.Println("")
}

func PrintMemStatuses(data []capture_processes.MemoryStatus) {
	table := table.New("PID", "Executable", "Memory (MB)")

	for _, mem_status := range data {
		table.AddRow(mem_status.Process.Pid, Truncate(mem_status.Process.Executable), KBToMB(mem_status.Memory))
	}

	table.Print()
}

func Truncate(s string) string {
	if len(s) > 50 {
		return s[:7] + "..." + s[len(s)-40:]
	}
	return s
}

func KBToMB(bytes uint64) uint64 {
	return bytes / uint64(1024)
}
