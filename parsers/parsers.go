package parsers

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func Parse(content []byte, path string) (map[string]interface{}, error) {
	var data map[string]interface{}
	ext := filepath.Ext(path)

	switch ext {
	case ".json":
		if err := json.Unmarshal(content, &data); err != nil {
			return nil, fmt.Errorf("invalid JSON: %w", err)
		}
	case ".yml", ".yaml":
		if err := yaml.Unmarshal(content, &data); err != nil {
			return nil, fmt.Errorf("invalid YAML: %w", err)
		}
	default:
		return nil, fmt.Errorf("unsupported format: %s (supported: .json, .yml, .yaml)", ext)
	}

	return data, nil
}
