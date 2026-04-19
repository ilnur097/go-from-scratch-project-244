package code

import (
    "os"
    "strings"
    "testing"
    
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestGenDiffNestedJSON(t *testing.T) {
    expectedBytes, err := os.ReadFile("testdata/fixtures/expected_nested.txt")
    require.NoError(t, err)
    expected := strings.TrimRight(string(expectedBytes), "\n\r\t ")
    
    result, err := GenDiff(
        "testdata/fixtures/filepath1_nested.json",
        "testdata/fixtures/filepath2_nested.json",
        "stylish",
    )
    require.NoError(t, err)
    result = strings.TrimRight(result, "\n\r\t ")
    
    assert.Equal(t, expected, result)
}

func TestGenDiffNestedYAML(t *testing.T) {
    expectedBytes, err := os.ReadFile("testdata/fixtures/expected_nested_yaml.txt")
    require.NoError(t, err)
    expected := strings.TrimRight(string(expectedBytes), "\n\r\t ")
    
    result, err := GenDiff(
        "testdata/fixtures/filepath1_nested.yml",
        "testdata/fixtures/filepath2_nested.yml",
        "stylish",
    )
    require.NoError(t, err)
    result = strings.TrimRight(result, "\n\r\t ")
    
    assert.Equal(t, expected, result)
}
