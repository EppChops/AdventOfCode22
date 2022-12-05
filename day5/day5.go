package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fileP := os.Args[1]

	data, _ := ioutil.ReadFile(fileP)

	split := strings.Split(string(data), "\n\n")
	stacks := parseStacks(split[0])

	for _, line := range strings.Split(split[1], "\n") {
		words := strings.Split(line, " ")
		amount, _ := strconv.Atoi(words[1])
		from, _ := strconv.Atoi(words[3])
		to, _ := strconv.Atoi(words[5])

		//Solution
		//for i := 1; i <= amount; i++{
		//move(&stacks, from, to)
		//}

		//Solution for part 2
		move2(&stacks, from, to, amount)
	}

	for _, strings := range stacks {
		if len(strings) > 0 {
			fmt.Print(strings[len(strings)-1])
		}
	}
	fmt.Println()
}

func move(stacks *[10][]string, src, dst int) {
	item := (*stacks)[src][len((*stacks)[src])-1]

	(*stacks)[src] = (*stacks)[src][:len((*stacks)[src])-1] //Pop

	(*stacks)[dst] = append((*stacks)[dst], item) //Add
}

func move2(stacks *[10][]string, src, dst, amount int) {
	items := (*stacks)[src][len((*stacks)[src])-amount:]

	(*stacks)[src] = (*stacks)[src][:len((*stacks)[src])-amount] //Pop n

	for _, item := range items {
		(*stacks)[dst] = append((*stacks)[dst], item) //Add
	}
}

func parseStacks(s string) [10][]string {
	lines := strings.Split(s, "\n")
	var stacks [10][]string
	zp := regexp.MustCompile(`    `)

	for i := len(lines) - 2; i >= 0; i-- {
		line := zp.ReplaceAllString(lines[i], " ")
		boxes := strings.Split(line, " ")

		for j, box := range boxes {
			if box != " " && box != "" && box != "  " {
				box = string(box[1])
				stacks[j+1] = append(stacks[j+1], box)
			}
		}
	}
	return stacks
}
