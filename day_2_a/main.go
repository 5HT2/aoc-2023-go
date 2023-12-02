package main

import (
	"encoding/csv"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	matchDice = regexp.MustCompile(`([0-9]+) (green|blue|red)`)
	matchGame = regexp.MustCompile(`(Game ([0-9]+))`)
)

func main() {
	b, err := os.ReadFile("day_2.txt")
	if err != nil {
		panic(err)
	}

	totalGameIDs := 0

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

			possible := true
			gameTitleMatch := matchGame.FindStringSubmatch(dice[0])
			gameID, err := strconv.Atoi(gameTitleMatch[2])
			if err != nil {
				log.Fatalf("failed to atoi %s: %s\n", gameTitleMatch, err)
			}

			for _, group := range dice {
				cubes := map[string]int{}
				matches := matchDice.FindAllStringSubmatch(group, -1)

				//log.Printf("%s\n", matches)
				for _, m1 := range matches {
					if len(m1) == 3 {
						numColorCube, err := strconv.Atoi(m1[1])
						if err != nil {
							log.Fatalf("couldn't atoi: %s: %s\n", m1[1], err)
						}

						// cubes[color] += total
						cubes[m1[2]] += numColorCube
					} else {
						log.Fatalf("len == %v: %s, '%s'\n", len(m1), group, m1)
					}

					if cubes["red"] > 12 || cubes["green"] > 13 || cubes["blue"] > 14 {
						possible = false
						//log.Printf("cubes: %s\n", cubes)
					} else {
						//log.Printf("cubes possible: %s\n", cubes)
					}

					//log.Printf("n0: %v / n1: %v / n2: %v / m1: %s\n", n0, n1, n2, m1)
				}
			}

			if possible {
				totalGameIDs += gameID
			}
		}

		log.Printf("%s\n", l)
	}

	log.Printf("output: %v\n", totalGameIDs)
}
