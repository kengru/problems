package advent

import (
	"bufio"
	"os"
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

	chars := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, r := range line {
			chars = append(chars, string(r))
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return chars
}
