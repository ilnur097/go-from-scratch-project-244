package code

import (
    "code/parsers"
    "fmt"
    "os"
    "sort"
)

type DiffNode struct {
    Key      string
    Status   string
    OldValue interface{}
    NewValue interface{}
    Children []DiffNode
}

func GenDiff(path1, path2, format string) (string, error) {
    data1, err := getData(path1)
    if err != nil {
        return "", err
    }
    data2, err := getData(path2)
    if err != nil {
        return "", err
    }

    diff := buildDiff(data1, data2)

    if format == "" || format == "stylish" {
        return formatStylish(diff), nil
    }

    return "", fmt.Errorf("unsupported format: %s", format)
}

func getData(path string) (map[string]interface{}, error) {
    content, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }
    return parsers.Parse(content, path)
}

func buildDiff(data1, data2 map[string]interface{}) []DiffNode {
    keys := getAllKeys(data1, data2)
    sort.Strings(keys)
    
    var diff []DiffNode
    
    for _, key := range keys {
        val1, ok1 := data1[key]
        val2, ok2 := data2[key]
        
        if !ok2 {
            diff = append(diff, DiffNode{
                Key:      key,
                Status:   "removed",
                OldValue: val1,
            })
        } else if !ok1 {
            diff = append(diff, DiffNode{
                Key:      key,
                Status:   "added",
                NewValue: val2,
            })
        } else if areMaps(val1, val2) {
 
            children := buildDiff(val1.(map[string]interface{}), val2.(map[string]interface{}))
            diff = append(diff, DiffNode{
                Key:      key,
                Status:   "nested",
                Children: children,
            })
        } else if val1 == val2 {
            diff = append(diff, DiffNode{
                Key:      key,
                Status:   "unchanged",
                OldValue: val1,
            })
        } else {
            diff = append(diff, DiffNode{
                Key:      key,
                Status:   "changed",
                OldValue: val1,
                NewValue: val2,
            })
        }
    }
    
    return diff
}

func getAllKeys(data1, data2 map[string]interface{}) []string {
    keys := make(map[string]bool)
    for k := range data1 {
        keys[k] = true
    }
    for k := range data2 {
        keys[k] = true
    }
    
    result := make([]string, 0, len(keys))
    for k := range keys {
        result = append(result, k)
    }
    return result
}

func areMaps(val1, val2 interface{}) bool {
    _, ok1 := val1.(map[string]interface{})
    _, ok2 := val2.(map[string]interface{})
    return ok1 && ok2
}
