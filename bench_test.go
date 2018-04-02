package main

import (
	"testing"
	"os"
	"bufio"
)

func BenchmarkPointArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		file, _ := os.Open("testdata.txt")
		defer file.Close()
		
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			RawPointInBoundingBox(scanner.Bytes())
		}
	}
}