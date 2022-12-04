package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Section struct {
	min int
	max int
}

func main() {

	fileP := os.Args[1]

	f, err := os.Open(fileP)
	check(err)

	scanner := bufio.NewScanner(f)

	count1 := 0
	count2 := 0
	for scanner.Scan() {
		sec1, sec2 := parsePairs(scanner.Text())

		if secContained(sec1, sec2) {
			count1++
		}

		if secOverlaps(sec1, sec2) {
			count2++
		}

	}

	fmt.Println(count2)
}

func secOverlaps(sec1, sec2 Section) bool {
	if (sec1.min <= sec2.min && sec1.max >= sec2.min) ||
		(sec1.min <= sec2.max && sec1.max >= sec2.max) ||
		(sec2.min <= sec1.min && sec2.max >= sec1.min) ||
		(sec2.min <= sec1.max && sec2.max >= sec1.min) {
		return true
	}
	return false
}

func secContained(sec1, sec2 Section) bool {
	if (sec1.min <= sec2.min && sec1.max >= sec2.max) ||
		(sec2.min <= sec1.min && sec2.max >= sec1.max) {
		return true
	}
	return false
}

func parsePairs(line string) (Section, Section) {
	pairs := strings.Split(line, ",")

	var sec1, sec2 Section
	var err error

	fstSec := strings.Split(pairs[0], "-")
	sec1.min, err = strconv.Atoi(fstSec[0])
	check(err)
	sec1.max, err = strconv.Atoi(fstSec[1])
	check(err)

	sndSec := strings.Split(pairs[1], "-")
	sec2.min, err = strconv.Atoi(sndSec[0])
	check(err)
	sec2.max, err = strconv.Atoi(sndSec[1])
	check(err)

	return sec1, sec2
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
