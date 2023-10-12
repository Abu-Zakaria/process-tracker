package report_preprocessor

import (
	"fmt"

	"github.com/Abu-Zakaria/process-tracker/internal/capture_processes"
	"github.com/Abu-Zakaria/process-tracker/internal/sorting"
	"github.com/rodaine/table"
)

type ProcessInfo struct {
	Pid        int
	Executable string
	MemorySize uint64
	Time       string
}

type Section struct {
	Title     string
	Processes []ProcessInfo
}

func TopTen(captures []capture_processes.Capture) {
	// sections := []Section{}
	fmt.Println("Preparing top ten processes...")

	sorted_captures := []capture_processes.Capture{}

	for _, capture := range captures {
		memory_statuses := capture.MemoryStatuses

		pivot := len(memory_statuses) - 1

		sorted_mem_statuses := sorting.QuickSortMems(memory_statuses, memory_statuses[pivot], pivot)
		sorted_mem_statuses = sorting.Reverse(sorted_mem_statuses)

		sorted_mem_statuses = takeFirstTen(sorted_mem_statuses)

		sorted_captures = append(sorted_captures, capture_processes.Capture{
			Time:                capture.Time,
			MemoryStatuses:      sorted_mem_statuses,
			TotalNumOfProcesses: len(memory_statuses),
		})
	}

	for _, sorted_capture := range sorted_captures {
		fmt.Println("List of processes (sorted)")

		fmt.Println("")
		fmt.Println("")

		printMemStatuses(sorted_capture.MemoryStatuses)
	}
}

func printMemStatuses(data []capture_processes.MemoryStatus) {
	table := table.New("PID", "Executable", "Memory (MB)")

	for _, mem_status := range data {
		table.AddRow(mem_status.Process.Pid, truncate(mem_status.Process.Executable), kBToMB(mem_status.Memory))
	}

	table.Print()
}

func takeFirstTen(data []capture_processes.MemoryStatus) []capture_processes.MemoryStatus {
	return data[:10]
}

func truncate(s string) string {
	if len(s) > 100 {
		return s[:40] + "..." + s[len(s)-40:]
	}
	return s
}

func kBToMB(bytes uint64) uint64 {
	return bytes / uint64(1024)
}
