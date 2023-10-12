package sorting

import (
	"github.com/Abu-Zakaria/process-tracker/internal/capture_processes"
)

func QuickSortMems(memory_statuses []capture_processes.MemoryStatus, pivot capture_processes.MemoryStatus, pivot_index int) []capture_processes.MemoryStatus {
	// take last item as pivot
	// start from 0th index, and go to bigger index until a number larger than pivot is found
	// take the larger item as first pointer
	// continue searching in the bigger indeces until a number smaller than pivot is found
	// take the smaller item as second pointer
	// swap the two pointers
	// if no smaller that pivot number is found, then swap pivot with first pointer
	// if no bigger number is found, take the index of (pivot - 1) as new pivot

	var first_pointer int = -1
	var second_pointer int = -1

	for i := 0; i < pivot_index; i++ {
		if first_pointer == -1 && memory_statuses[i].Memory > pivot.Memory {
			first_pointer = i
		}

		if first_pointer != -1 && second_pointer == -1 && memory_statuses[i].Memory < pivot.Memory {
			second_pointer = i
			break
		}
	}

	if first_pointer != -1 && second_pointer != -1 {
		memory_statuses[first_pointer], memory_statuses[second_pointer] = memory_statuses[second_pointer], memory_statuses[first_pointer]
	} else if first_pointer != -1 && second_pointer == -1 {
		memory_statuses[first_pointer], memory_statuses[pivot_index] = memory_statuses[pivot_index], memory_statuses[first_pointer]
	} else if pivot_index = pivot_index - 1; pivot_index == 0 {
		return memory_statuses
	}

	return QuickSortMems(memory_statuses, memory_statuses[pivot_index], pivot_index)
}

func Reverse(memory_statuses []capture_processes.MemoryStatus) []capture_processes.MemoryStatus {
	reversed := []capture_processes.MemoryStatus{}

	for i := len(memory_statuses) - 1; i >= 0; i-- {
		reversed = append(reversed, memory_statuses[i])
	}

	return reversed
}
