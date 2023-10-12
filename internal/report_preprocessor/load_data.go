package report_preprocessor

import (
	"log"

	"github.com/Abu-Zakaria/process-tracker/internal/capture_processes"
	"github.com/Abu-Zakaria/process-tracker/pkg/json_data_handler"
)

func GetData(file_path string) []capture_processes.Capture {
	captures := []capture_processes.Capture{}

	err := json_data_handler.ReadJSON(file_path, &captures)
	if err != nil {
		log.Println("Couldn't read", file_path, ". Error Message -", err)
	}

	return captures
}
