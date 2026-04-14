package code

import (
	"os"
	"testing"
	"github.com/urfave/cli/v3"
)

func TestGenDiff(t *testing.T) {
	expected, _ := os.ReadFile("testdata/fixtures/expected_result.txt")
	
	result, err := GenDiff(
		"testdata/fixtures/file1.json",
		"testdata/fixtures/file2.json",
		"stylish",
	)
	
	assert.NoError(t, err)
	assert.Equal(t, string(expected), result)
}
