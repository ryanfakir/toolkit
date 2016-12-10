package utility

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(filename string) map[string]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("file open occur error")
	}
	dict := make(map[string]int, 0)
	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		line := input.Text()
		dict[line]++
	}
	return dict
}
