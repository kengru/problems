package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	year := flag.Int("year", 0, "Year of advent of code problem (e.g. 2025)")
	day := flag.Int("day", 0, "Day of advent of code problem (e.g. 5)")
	part := flag.Int("part", 0, "Part of the problem (1 or 2)")
	flag.Parse()

	if *year == 0 || *day == 0 || *part == 0 {
		fmt.Println("Usage: go run . --year YYYY --day DD --part P")
		fmt.Println("Example: go run . --year 2025 --day 5 --part 1")
		return
	}

	if *part < 1 || *part > 2 {
		log.Fatal("Part must be 1 or 2")
	}

	err := RunAdventProblem(*year, *day, *part)
	if err != nil {
		log.Fatal(err)
	}
}
