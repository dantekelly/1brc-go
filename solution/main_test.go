package main

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}

func TestCalculateAverages(t *testing.T) {
	expected, err := os.ReadFile("../averages.txt")
	if err != nil {
		t.Errorf("Could not read expected output file: %v", err)
	}

	output := captureOutput(calculateAverages)
	if !strings.Contains(output, "Result:") {
		t.Errorf("calculateAverages() did not complete successfully")
	}

	result := strings.Split(output, ": ")[1]
	if string(expected) != result {
		t.Errorf("Expected %s, got %s", string(expected), result)
	}
}

func BenchmarkReadFile(b *testing.B) {
	var measurements map[string]Measurement
	for i := 0; i < b.N; i++ {
		readFile(&measurements)
	}
}

func BenchmarkCalculateAverage(b *testing.B) {
	var measurements map[string]Measurement
	readFile(&measurements)
	for i := 0; i < b.N; i++ {
		calculateAverage(measurements)
	}
}
