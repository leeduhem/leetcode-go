package p195

import (
	"bufio"
	"log"
	"os"
)

func nthLine(n int, file *os.File) string {
	if n <= 0 {
		return ""
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		n--
		if n == 0 {
			return scanner.Text()
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return ""
}

func tenthLine(name string) string {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	return nthLine(10, file)
}
