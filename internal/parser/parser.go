package code

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)


func ParseFile(path string) (map[string]interface{}, error) {
	
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("could not read file %s: %w", path, err)
	}

	var data map[string]interface{}
	
	if err := json.Unmarshal(content, &data); err != nil {
		return nil, fmt.Errorf("could not parse JSON from %s: %w", path, err)
	}

	return data, nil
}
