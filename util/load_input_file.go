package util

import (
	"bufio"
	"log"
	"os"
)

func LoadInputFile(s string) []string {
	err := error(nil)
	f, err := os.Open(s)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
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
