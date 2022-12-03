package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	fileP := os.Args[1]

	f, err := os.Open(fileP)
	check(err)

	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanLines)

	var duplicates []rune
	priorities := make(map[rune]int)

	i := 1
	for r := 'a'; r <= 'z'; r++ {
		priorities[r] = i
		priorities[unicode.ToUpper(r)] = i + 26
		i++
	}

	for scanner.Scan() {
		rucksack := []rune(scanner.Text())
		mid := len(rucksack) / 2

		fstCompartment := rucksack[:mid]
		sndCompartment := rucksack[mid:]

		for _, cFst := range fstCompartment {
			if contains(sndCompartment, cFst) {
				duplicates = append(duplicates, cFst)
				break
			}
		}
	}

	f.Close()

	count := 0
	for _, dup := range duplicates {
		count += priorities[dup]
	}
	fmt.Println(count)
}

func contains(arr []rune, item rune) bool {
	for _, r := range arr {
		if item == r {
			return true
		}
	}
	return false
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
