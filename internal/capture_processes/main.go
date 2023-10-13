package capture_processes

import (
	"fmt"
	"time"

	"github.com/Abu-Zakaria/process-tracker/pkg/json_data_handler"
)

var (
	OS string
)

type Capture struct {
	Time                string         `json:"time"`
	MemoryStatuses      []MemoryStatus `json:"memory_statuses"`
	TotalNumOfProcesses int            `json:"total_num_of_processes"`
}

func Run(mem_data_file_path string) {
	for {
		SaveMemData(mem_data_file_path)
		time.Sleep(30 * time.Second)
	}
}

func SaveMemData(mem_data_file_path string) {
	mems := CaptureMems()

	capture := Capture{
		Time:                time.Now().Format(time.RFC3339),
		MemoryStatuses:      mems,
		TotalNumOfProcesses: len(mems),
	}

	old_data := []Capture{}
	captures := []Capture{}

	err := json_data_handler.ReadJSON(mem_data_file_path, &old_data)
	if err != nil {
		fmt.Println("Couldn't find any existing data file. Creating new data file...")
	} else {
		captures = old_data
	}

	captures = append(captures, capture)

	err = json_data_handler.SaveJSON(captures, mem_data_file_path)
	if err != nil {
		fmt.Println("Couldn't save process memory data to", mem_data_file_path)
	}

	fmt.Println("Data saved to mem_data.json")
}

func CaptureMems() []MemoryStatus {
	processes := GetAllProcesses()
	memory_stats := []MemoryStatus{}

	for _, process := range processes {
		memory_stats = append(memory_stats, GetProcessMemory(process))
	}

	return memory_stats
}
