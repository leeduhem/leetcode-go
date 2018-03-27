package p193

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

func validPhoneNumbers(name string) []string {
	const phoneNumberPat = `^(?:\(\d{3}\)\s|\d{3}-)\d{3}-\d{4}$`

	phoneNumberRE, err := regexp.Compile(phoneNumberPat)
	if err != nil {
		log.Fatal(err)
	}

	phoneNumbers := make([]string, 0)

	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if phoneNumberRE.MatchString(line) {
			phoneNumbers = append(phoneNumbers, line)
		}
	}

	return phoneNumbers
}
