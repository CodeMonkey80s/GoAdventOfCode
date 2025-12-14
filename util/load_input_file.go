package util

import (
	"bufio"
	"log"
	"os"
)

func LoadInputFile(s string) []string {
	f, err := os.Open(s)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}(f)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
