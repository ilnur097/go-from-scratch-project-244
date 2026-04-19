package code

import (
    "os"
    "strings"
    "testing"
"github.com/stretchr/testify/assert"
    
)

func TestGenDiff(t *testing.T) {
    expectedBytes, err := os.ReadFile("testdata/fixtures/expected_result.txt")
    if err != nil {
        t.Fatalf("failed to read expected result: %v", err)
    }
    expected := strings.TrimRight(string(expectedBytes), "\n\r\t ")

    result, err := GenDiff(
        "testdata/fixtures/filepath1.json",
        "testdata/fixtures/filepath2.json",
        "stylish",
    )
    if err != nil {
        t.Fatalf("GenDiff returned error: %v", err)
    }
    result = strings.TrimRight(result, "\n\r\t ")

    if result != expected {
        t.Errorf("expected:\n%s\ngot:\n%s", expected, result)
    }
}
func TestGenDiffYAML(t *testing.T) {
	file1 := "testdata/fixtures/file1.yml"
	file2 := "testdata/fixtures/file2.yml"
	expectedPath := "testdata/fixtures/expected_flat.txt"

	expectedBytes, _ := os.ReadFile(expectedPath)
	expected := strings.TrimSpace(string(expectedBytes))

	result, err := GenDiff(file1, file2, "stylish")

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}
