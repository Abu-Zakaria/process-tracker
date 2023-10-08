package json_data_handler

import (
	"encoding/json"
	"os"
)

func SaveJSON(data any, file_name string) error {
	data_bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(file_name, data_bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadJSON(file_name string, unmarshal_obj any) error {
	data, err := os.ReadFile(file_name)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, unmarshal_obj)
	if err != nil {
		return err
	}

	return nil
}
