package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileP := os.Args[1]

	//Signal strength = cyclenr * xValue during every 20th cycle

	f, _ := os.Open(fileP)
	scanner := bufio.NewScanner(f)

	cycle := 0
	regX := 1
	//nr := 0
	signalStrengths := make([]int, 0)

	var crt [6][40]string

	for scanner.Scan() {
		//fmt.Println(nr, cycle)
		//nr++
		line := scanner.Text()
		if line == "noop" {
			cycle++
			addSprite(regX, cycle, &crt)
			if cycle%20 == 0 {
				fmt.Println("sig str noop", cycle, regX, (cycle * regX))
				signalStrengths = append(signalStrengths, calcSignalStrength(cycle, regX))
			}
		} else {
			cycle++
			addSprite(regX, cycle, &crt)
			if cycle%20 == 0 {
				fmt.Println("sig str op 1", cycle, regX, (cycle * regX))
				signalStrengths = append(signalStrengths, calcSignalStrength(cycle, regX))
			}

			cycle++
			addSprite(regX, cycle, &crt)
			if cycle%20 == 0 {
				fmt.Println("sig str op 2", cycle, regX, (cycle * regX))
				signalStrengths = append(signalStrengths, calcSignalStrength(cycle, regX))
			}
			add, _ := strconv.Atoi(strings.Split(line, " ")[1])
			regX += add
			
		}
	}
	fmt.Println(signalStrengths)
	sum := 0
	for i := range signalStrengths {
		if i %2 == 0 {
			sum += signalStrengths[i]
		}
	}
	fmt.Println(sum)
	prettyPrint(&crt)
}

func prettyPrint(crt *[6][40]string) {
	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			fmt.Print(crt[i][j])
		}
		fmt.Print("\n")
	}
}

func addSprite(regX int, cycle int, crt *[6][40]string) {
	var col, row int
	if cycle >= 1 && cycle <= 40{
		col = cycle - 1
		row = 0
	} else if cycle >= 41 && cycle <= 80 {
		col = cycle - 41
		row = 1
	} else if cycle >= 81 && cycle <= 120 {
		col = cycle - 81
		row = 2
	} else if cycle >= 121 && cycle <= 160 {
		col = cycle - 121
		row = 3
	} else if cycle >= 161 && cycle <= 200 {
		col = cycle - 161
		row = 4
	} else if cycle >= 201 && cycle <= 240 {
		col = cycle - 201
		row = 5
	}

	if regX >= col - 1 && regX <= col + 1 {
		crt[row][col] = "#"
	} else {
		crt[row][col] = "."
	}
} 

func calcSignalStrength(cycle, regX int) int {
	return cycle * regX
}
