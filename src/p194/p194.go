package p194

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func linesOf(name string) []string {
	lines := make([]string, 0)

	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func split(s string) []string {
	return strings.Split(s, " ")
}

func transpose(m [][]string) [][]string {
	mp := make([][]string, len(m[0]))

	for j, _ := range m[0] {
		mp[j] = make([]string, len(m))
		for i, row := range m {
			mp[j][i] = row[j]
		}
	}

	return mp
}

func transposeFile(name string) []string {
	lines := linesOf(name)
	strMatrix := make([][]string, len(lines))
	for i, line := range lines {
		strMatrix[i] = split(line)
	}

	transposedFile := make([]string, len(strMatrix[0]))
	for i, row := range transpose(strMatrix) {
		transposedFile[i] = strings.Join(row, " ")
	}
	return transposedFile
}
