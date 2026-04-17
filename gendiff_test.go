package code

import (
    "os"
    "strings"
    "testing"
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
