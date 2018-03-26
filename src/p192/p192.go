package p192

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func wordFrequency(name string) []string {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}

	wordFreq := make(map[string]int)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		wordFreq[word]++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	words := make([]string, len(wordFreq))
	idx := 0
	for word, _ := range wordFreq {
		words[idx] = word
		idx++
	}

	sort.Slice(words, func(i, j int) bool {
		return wordFreq[words[i]] >= wordFreq[words[j]]
	})

	freqs := make([]string, len(words))
	for idx, word := range words {
		freqs[idx] = word + " " + strconv.Itoa(wordFreq[word])
	}

	return freqs
}
