package code

import (
	
	"fmt"
	"os"
	"sort"
	"strings"
        "code/parsers"
)


func GenDiff(path1, path2, format string) (string, error) {
	
	data1, err := getData(path1)
	if err != nil {
		return "", err
	}
	data2, err := getData(path2)
	if err != nil {
		return "", err
	}

	
	keysMap := make(map[string]bool)
	for k := range data1 {
		keysMap[k] = true
	}
	for k := range data2 {
		keysMap[k] = true
	}

	keys := make([]string, 0, len(keysMap))
	for k := range keysMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	
	var lines []string
	lines = append(lines, "{")

	for _, key := range keys {
		val1, ok1 := data1[key]
		val2, ok2 := data2[key]

		if !ok2 {
			lines = append(lines, fmt.Sprintf("  - %s: %v", key, val1))
		} else if !ok1 {
			lines = append(lines, fmt.Sprintf("  + %s: %v", key, val2))
		} else if val1 == val2 {
			lines = append(lines, fmt.Sprintf("    %s: %v", key, val1))
		} else {
			lines = append(lines, fmt.Sprintf("  - %s: %v", key, val1))
			lines = append(lines, fmt.Sprintf("  + %s: %v", key, val2))
		}
	}

	lines = append(lines, "}")
	return strings.Join(lines, "\n"), nil
}


func getData(path string) (map[string]interface{}, error) {
    content, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }
    return parsers.Parse(content, path)
}
