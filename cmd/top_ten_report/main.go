package main

import "github.com/Abu-Zakaria/process-tracker/internal/report_preprocessor"

const (
	memory_data_file_path = "./mem_data.json"
)

func main() {
	data := report_preprocessor.GetData(memory_data_file_path)

	report_preprocessor.TopTen(data)
}
