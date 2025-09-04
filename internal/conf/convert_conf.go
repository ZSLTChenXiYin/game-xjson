package conf

import (
	"encoding/json"
	"io"
	"os"
)

type ConvertConfig struct {
	Table        string  `json:"table"`
	Output       *string `json:"output"`
	Sheet        *string `json:"sheet"`
	Types        *bool   `json:"types"`
	Descriptions *bool   `json:"descriptions"`
	Pretty       *bool   `json:"pretty"`
	IncrementID  *bool   `json:"increment_id"`
}

func LoadConvertConfig(path string) (*ConvertConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	config_json, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	convert_config := &ConvertConfig{}

	err = json.Unmarshal(config_json, convert_config)
	if err != nil {
		return nil, err
	}

	return convert_config, nil
}
