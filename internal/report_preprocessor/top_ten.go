package report_preprocessor

import (
	"fmt"

	"github.com/Abu-Zakaria/process-tracker/internal/capture_processes"
	"github.com/Abu-Zakaria/process-tracker/internal/sorting"
)

func TopTen(captures []capture_processes.Capture) {
	fmt.Println("Preparing top ten processes...")
	fmt.Println("")

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

	fmt.Println("List of processes (sorted)")

	for _, sorted_capture := range sorted_captures {
		RenderCapture(sorted_capture)
	}
}

func takeFirstTen(data []capture_processes.MemoryStatus) []capture_processes.MemoryStatus {
	return data[:10]
}
