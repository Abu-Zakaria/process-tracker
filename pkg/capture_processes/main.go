package capture_processes

import (
	"log"
	"time"

	"github.com/Abu-Zakaria/process-tracker/pkg/json_data_handler"
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
		log.Println("Couldn't read mem_data.json before saving new data to it! Error Message -", err)
	} else {
		captures = old_data
	}

	captures = append(captures, capture)

	err = json_data_handler.SaveJSON(captures, "mem_data.json")
	if err != nil {
		log.Fatal("Couldn't save data to mem_data.json: ", err)
	}

	log.Println("Data saved to mem_data.json")
}

func CaptureMems() []MemoryStatus {
	processes := GetAllProcesses()
	memory_stats := []MemoryStatus{}

	for _, process := range processes {
		memory_stats = append(memory_stats, GetProcessMemory(process))
	}

	return memory_stats
}
