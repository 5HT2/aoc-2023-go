package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var (
	words    = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	wordsNew = map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
)

func main() {
	start := time.Now()
	b, err := os.ReadFile("day_1.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")
	total := 0

	log.Printf("%v\n", len(lines))
	for _, s := range lines {
		if len(s) == 0 {
			continue
		}
		//
		//first := ""
		//last := ""
		//
		//// 1

		// 2
		firstIndex := -1
		lastIndex := -1
		firstIntIndex := -1
		lastIntIndex := -1

		first := ""
		last := ""
		firstDigit := ""
		lastDigit := ""

		for word, wordInt := range wordsNew {
			if i := strings.Index(s, word); i != -1 {
				if i < firstIndex || firstIndex == -1 {
					first = wordInt
					firstIndex = i
				}
			}

			if i := getLastIndex(s, word); i != -1 {
				if i >= lastIndex {
					last = wordInt
					lastIndex = i
				}
			}
		}

		for n, r := range s {
			if unicode.IsDigit(r) {
				if firstIndex == -1 || n == 0 || n < firstIndex {
					first = string(r)
					firstDigit = first
					firstIntIndex = n
				}
				break
			}
		}

		runes := []rune(s)
		for i := len(runes) - 1; i >= 0; i-- {
			if unicode.IsDigit(runes[i]) {
				if lastIndex == -1 || i > lastIndex {
					last = string(runes[i])
					lastDigit = last
					lastIntIndex = i
				}
				break
			}
		}

		//log.Printf("%s: %s / %s\n", s, first, last)

		num, err := strconv.Atoi(first + last)
		if err != nil {
			log.Fatalf("couldn't atoi: %s / %s: %s\n", first, last, err)
		}

		total += num
		log.Printf("%s: %v (%s / %s / %s / %s / %v / %v / %v / %v)\n", s, num, first, last, firstDigit, lastDigit, firstIndex, lastIndex, firstIntIndex, lastIntIndex)
	}

	stop := time.Now()
	diff := float32(stop.UnixNano() - start.UnixNano())
	log.Printf("final answer is %v, took %.2fms\n",
		total, diff/1000000,
	)
}

func getLastIndex(s, substr string) int {
	index := -1
	for {
		next := strings.Index(s, substr)
		if next == -1 {
			return index
		}
		index += next + len(substr)
		s = s[next+len(substr):]
	}
}
