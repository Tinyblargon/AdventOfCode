package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"testing"
)

func Benchmark_scan(b *testing.B) {
	content, err := os.ReadFile("test-data")
	if err != nil {
		log.Fatal(err)
	}
	var characters [][]rune
	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	for scanner.Scan() {
		characters = append(characters, []rune(scanner.Text()))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		scan(characters)
	}
}
