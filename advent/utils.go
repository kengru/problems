package advent

import (
	"bufio"
	"os"
	"strings"
)

func GetLineFromFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines[0]
}

func GetLinesFromFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func GetMatrixFromFile(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := [][]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		characters := []string{}
		for _, l := range scanner.Text() {
			characters = append(characters, string(l))
		}
		lines = append(lines, characters)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func GetCharactersFromFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return strings.Split(line, "")
}

func GetSeparatedFromFile(path string, trim bool) [][]string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := [][]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		characters := []string{}
		numbers := strings.Split(scanner.Text(), " ")
		for _, l := range numbers {
			if trim {
				n := strings.TrimSpace(l)
				if n != "" {
					characters = append(characters, n)
				}
			} else {
				characters = append(characters, l)
			}
		}
		lines = append(lines, characters)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}
