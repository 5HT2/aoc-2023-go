package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	matchDice = regexp.MustCompile(`([0-9]+) (red|green|blue)`)
)

type Cubes map[string]int

func (c Cubes) String() string {
	return fmt.Sprintf("[r: %v, g: %v, b: %v]", c["red"], c["green"], c["blue"])
}

func main() {
	b, err := os.ReadFile("day2.txt")
	if err != nil {
		panic(err)
	}

	totalPowers := 0

	s := strings.Split(string(b), "\n")
	for _, l := range s { // this is the worst fucking parsing i've written in 2023. it's so slow.
		if len(s) == 0 {
			return
		}

		reader := csv.NewReader(strings.NewReader(l))
		reader.Comma = ';'

		subGames, err := reader.ReadAll()
		if err != nil {
			log.Fatalf("%s: %s\n", l, err)
			//continue
		}

		for _, dice := range subGames {
			if len(dice) == 0 {
				continue
			}

			cubes := Cubes{}
			for _, group := range dice {
				matches := matchDice.FindAllStringSubmatch(group, -1)

				//log.Printf("%s\n", matches)
				for _, m1 := range matches {
					if len(m1) == 3 {
						numColorCube, err := strconv.Atoi(m1[1])
						if err != nil {
							log.Fatalf("couldn't atoi: %s: %s\n", m1[1], err)
						}

						if numColorCube > cubes[m1[2]] {
							cubes[m1[2]] = numColorCube
						}
					} else {
						log.Fatalf("len == %v: %s, '%s'\n", len(m1), group, m1)
					}

					//if cubes["red"] > 12 || cubes["green"] > 13 || cubes["blue"] > 14 {
					//	possible = false
					//	log.Printf("cubes: %s\n", cubes)
					//} else {
					//	log.Printf("cubes possible: %s\n", cubes)
					//}

					//log.Printf("n0: %v / n1: %v / n2: %v / m1: %s\n", n0, n1, n2, m1)
				}

			}

			log.Printf("cubes: %s\n", cubes)
			totalPowers += cubes["red"] * cubes["green"] * cubes["blue"]
		}

		log.Printf("%s\n", l)
	}

	log.Printf("output: %v\n", totalPowers)
}
