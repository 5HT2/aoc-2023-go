package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func main() {
	start := time.Now()
	b, err := os.ReadFile("day_1.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")
	total := 0

	for _, s := range lines {
		if len(s) == 0 {
			continue
		}

		first := ""
		last := ""
		for _, r := range s {
			if unicode.IsDigit(r) {
				first = string(r)
				break
			}
		}

		runes := []rune(s)
		for i := len(runes) - 1; i >= 0; i-- {
			if unicode.IsDigit(runes[i]) {
				last = string(runes[i])
				break
			}
		}

		num, err := strconv.Atoi(first + last)
		if err != nil {
			log.Fatalf("couldn't atoi: %s / %s: %s\n", first, last, err)
		}

		total += num
		//log.Printf("%s: %v (%s / %s)\n", s, num, first, last)
	}

	stop := time.Now()
	diff := float32(stop.UnixNano() - start.UnixNano())
	log.Printf("final answer is %v, took %.2fms\n",
		total, diff/1000000,
	)
}
