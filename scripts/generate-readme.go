package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	progressTablePlaceholder = "[progress-table-goes-here]"
	dirSolutions             = "./solutions/"
	templateFile             = "README-template.md"
	outputFile               = "README.md"
)

func main() {

	output := []string{
		"|  #   | Days                                                                                                      |",
		"|------|-----------------------------------------------------------------------------------------------------------|",
	}

	realPath, _ := filepath.Abs(dirSolutions)
	yearItems, err := os.ReadDir(realPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range yearItems {
		if f.IsDir() {
			dirName := f.Name()
			if len(dirName) == 4 {
				progressLine := progressForYear(dirName)
				yearDir, _ := filepath.Abs(dirSolutions + dirName)
				solutionItems, err := os.ReadDir(yearDir)
				if err != nil {
					log.Fatal(err)
				}
				for _, f2 := range solutionItems {
					if f2.IsDir() {
						dirname := f2.Name()
						if !strings.Contains(dirname, "day") {
							continue
						}

						parts := strings.Split(dirname, "_")
						s := strings.ReplaceAll(parts[0], "day", "")
						n, _ := strconv.Atoi(s)
						if n > 0 {
							fmt.Printf("year: %s, checking day: %d\n", yearDir, n)
						}

						if parts[1] == "part1" {
							idx := 8 + n
							progressLine[idx] = '✅'
						}
						if parts[1] == "part2" {
							idx := 38 + n
							progressLine[idx] = '✅'
						}

						// filenamePart1 := fmt.Sprintf("%s/%s/part1.go", yearDir, dirname)
						// fmt.Printf("checking file: %q\n", filenamePart1)
						// if _, err := os.Stat(filenamePart1); !errors.Is(err, fs.ErrNotExist) {
						// 	idx := 8 + n
						// 	progressLine[idx] = '✅'
						// }
						//
						// filenamePart2 := fmt.Sprintf("%s/%s/part2.go", yearDir, dirname)
						// fmt.Printf("checking file: %q\n", filenamePart2)
						// if _, err := os.Stat(filenamePart2); !errors.Is(err, fs.ErrNotExist) {
						// 	idx := 38 + n
						// 	progressLine[idx] = '✅'
						// }
					}
				}
				output = append(output, string(progressLine))
			}
		}
	}

	updateReadme(templateFile, outputFile, output)

}

func progressForYear(year string) []rune {
	output := fmt.Sprintf("| %s | 🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲<br/>🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲 |", year)
	return []rune(output)
}

func updateReadme(templateFile string, outputFile string, output []string) {
	b, err := os.ReadFile(templateFile)
	if err != nil {
		log.Fatal(err)
	}

	b = []byte(strings.ReplaceAll(string(b), progressTablePlaceholder, strings.Join(output, "\n")))
	_, err = os.Stdout.Write(b)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(outputFile, b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
