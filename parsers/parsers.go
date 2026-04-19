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
			return nil, err
		}
	case ".yml", ".yaml":
		if err := yaml.Unmarshal(content, &data); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown format: %s", ext)
	}

	return data, nil
}
