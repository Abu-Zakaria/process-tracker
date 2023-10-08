package capture_processes

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type Capture struct {
	Time                string         `json:"time"`
	MemoryStatuses      []MemoryStatus `json:"memory_statuses"`
	TotalNumOfProcesses int            `json:"total_num_of_processes"`
}

func Run() {
	for {
		SaveMemData()
		time.Sleep(30 * time.Second)
	}
}

func SaveMemData() {
	mems := CaptureMems()

	capture := Capture{
		Time:                time.Now().Format(time.RFC3339),
		MemoryStatuses:      mems,
		TotalNumOfProcesses: len(mems),
	}

	data, err := json.Marshal(capture)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("mem_data.json", data, 0644)
	if err != nil {
		log.Println("Not able to save data to mem_data.json")
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
