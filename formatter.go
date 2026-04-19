package code

import (
    "fmt"
    "sort"
    "strings"
)

func formatStylish(diff []DiffNode) string {
    return formatNode(diff, 0)
}

func formatNode(nodes []DiffNode, depth int) string {
    if len(nodes) == 0 {
        return "{}"
    }
    
    indent := strings.Repeat("  ", depth)
    result := "{\n"
    
    for _, node := range nodes {
        switch node.Status {
        case "nested":
            result += fmt.Sprintf("%s  %s: %s", indent, node.Key, formatNode(node.Children, depth+1))
        case "added":
            result += fmt.Sprintf("%s+ %s: %s\n", indent, node.Key, formatValue(node.NewValue, depth+1))
        case "removed":
            result += fmt.Sprintf("%s- %s: %s\n", indent, node.Key, formatValue(node.OldValue, depth+1))
        case "changed":
            result += fmt.Sprintf("%s- %s: %s\n", indent, node.Key, formatValue(node.OldValue, depth+1))
            result += fmt.Sprintf("%s+ %s: %s\n", indent, node.Key, formatValue(node.NewValue, depth+1))
        case "unchanged":
            result += fmt.Sprintf("%s  %s: %s\n", indent, node.Key, formatValue(node.OldValue, depth+1))
        }
    }
    
    result += indent + "}\n"
    return result
}

func formatValue(value interface{}, depth int) string {
    if value == nil {
        return "null"
    }
    
 
    if m, ok := value.(map[string]interface{}); ok {
 
        var nodes []DiffNode
        for k, v := range m {
            nodes = append(nodes, DiffNode{
                Key:      k,
                Status:   "unchanged",
                OldValue: v,
            })
        }
 
        sort.Slice(nodes, func(i, j int) bool {
            return nodes[i].Key < nodes[j].Key
        })
        return formatNode(nodes, depth)
    }
    
 
    switch v := value.(type) {
    case string:
        return v
    case float64:
        if v == float64(int(v)) {
            return fmt.Sprintf("%d", int(v))
        }
        return fmt.Sprintf("%g", v)
    case bool:
        return fmt.Sprintf("%t", v)
    default:
        return fmt.Sprintf("%v", v)
    }
}
