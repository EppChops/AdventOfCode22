package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func main() {
	fileP := os.Args[1]

	f, _ := os.Open(fileP)

	scanner := bufio.NewScanner(f)

	tailPositions := make(map[Point]int)

	var tail, head Point
	tail.x = 0
	tail.y = 0

	head.x = 0
	head.y = 0

	tailPositions[tail]++

	var rope [10]*Point
	for i := range rope{
		rope[i] = &Point{0,0}
	}

	for scanner.Scan(){
		motions := strings.Split(scanner.Text(), " ")
		cmd := motions[0]
		amount, _ := strconv.Atoi(motions[1])
		move2(&head, &rope, cmd, amount, &tailPositions)
	}

	fmt.Println(len(tailPositions))
}

func move2(head *Point, rope *[10]*Point , direction string, amount int, tailPostitions * map[Point]int){
	fmt.Println(direction, amount)
	switch direction {
	case "R":
		for i := 0; i < amount; i++{
			(*rope)[0].x++
			for j := 0; j < len(rope) -1; j++ {

				if(!isAdjacent(rope[j], rope[j+1])) {
					moveTail(rope[j], rope[j+1])
					if j == 8 {
						(*tailPostitions)[*rope[j+1]]++
					}
				}
			}
			
		}
	case "U":
		for i := 0; i < amount; i++{
			(*rope)[0].y++
			for j := 0; j < len(rope)-1; j++ {
				if(!isAdjacent(rope[j], rope[j+1])) {
					moveTail(rope[j], rope[j+1])
					if j == 8 {
						(*tailPostitions)[*rope[j+1]]++
					}
				}
			}
		}
	case "L":
		for i := 0; i < amount; i++{
			(*rope)[0].x--
			for j := 0; j < len(rope)-1; j++ {
				if(!isAdjacent(rope[j], rope[j+1])) {
					
					moveTail(rope[j], rope[j+1])
					if j == 8 {
						(*tailPostitions)[*rope[j+1]]++
					}
				}
			}
		}
	case "D":
		for i := 0; i < amount; i++{
			(*rope)[0].y--
			for j := 0; j < len(rope)-1; j++ {
				if(!isAdjacent(rope[j], rope[j+1])) {
					
					moveTail(rope[j], rope[j+1])
					if j == 8 {
						(*tailPostitions)[*rope[j+1]]++
					}
				}
			}
		}
	}

	
}

func move(head, tail *Point, direction string, amount int, tailPostitions * map[Point]int){
	switch direction {
	case "R":
		for i := 0; i < amount; i++{
			head.x++
			if(!isAdjacent(head, tail)) {
				moveTail(head, tail)
				(*tailPostitions)[*tail]++
			}
		}
	case "U":
		for i := 0; i < amount; i++{
			head.y++
			if !isAdjacent(head, tail) {
				moveTail(head, tail)
				(*tailPostitions)[*tail]++
			}
		}
	case "L":
		for i := 0; i < amount; i++{
			head.x--
			if !isAdjacent(head, tail) {
				moveTail(head, tail)
				(*tailPostitions)[*tail]++
			}
		}
	case "D":
		for i := 0; i < amount; i++{
			head.y--
			if !isAdjacent(head, tail) {
				moveTail(head, tail)
				(*tailPostitions)[*tail]++
			}
		}
	}

}

func moveTail(head, tail *Point) {
	if head.x == tail.x {
		if head.y > tail.y {
			tail.y++
		} else {
			tail.y--
		}
	} else if head.y == tail.y {
		if head.x > tail.x {
			tail.x++
		}else {
			tail.x--
		}
	} else {
			if head.x > tail.x {
				tail.x++
			}else {
				tail.x--
			}
			if head.y > tail.y {
					tail.y++
			} else {
				tail.y--
		}
	}
}

func isAdjacent(head, tail *Point) bool {
	dx := Abs(head.x - tail.x)
	dy := Abs(head.y - tail.y)
	adj := (dx + dy == 1) || (dx == dy && dx + dy == 2) || 
		(head.x == tail.x && head.y == tail.y)
	
	return adj
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}