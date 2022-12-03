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

	var badges []rune
	priorities := make(map[rune]int)

	i := 1
	for r := 'a'; r <= 'z'; r++ {
		priorities[r] = i
		priorities[unicode.ToUpper(r)] = i + 26	
		i++
	}

	i = 0
	var lines [][]rune
	for scanner.Scan(){
		rucksack := []rune(scanner.Text())
		lines = append(lines, rucksack)
	}

	for i = 0; i < len(lines)-2; i++{
		r := compareSacks(lines[i], lines[i+1], lines[i+2])
		i += 2
		if r != 0 {
			badges = append(badges, r)
		}
	}
	f.Close()

	count := 0
	for _, dup := range badges {
		count += priorities[dup]
	}
	fmt.Println(count)
}

func compareSacks(sack1, sack2, sack3 []rune) rune {
	for _, c := range sack1 {
		if contains(sack2, c) && contains(sack3, c) {
				return c
		}
	}
	return 0
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