package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Monkey struct {
	id          int
	test        int
	operationOp string
	operation   int
	trueThrow   int
	falseThrow  int
	inspections int
	items       []*Item
}

type Item struct {
	worryLvl *big.Int
}

func main() {
	fileP := os.Args[1]

	data, _ := ioutil.ReadFile(fileP)
	ms := strings.Split(string(data), "\n\n")

	var monkeys []*Monkey
	for _, m := range ms {
		monkey := parseMonkey(m)
		monkeys = append(monkeys, &monkey)
	}

	for rounds := 0; rounds < 1000; rounds++ {
		for _, m := range monkeys {
			//for all items
			//inspect
			for _, item := range m.items {
				worrylvl := m.inspect(item)
				//worrylvl = worrylvl / 3
				item.worryLvl = worrylvl
				//fmt.Println(worrylvl, item.worryLvl)
				throwTo := m.testThrow(worrylvl)
				for _, x := range monkeys {
					if x.id == throwTo {
						m.throw(x, item)
						break
					}
				}
			}
			m.items = []*Item{}
		}
	}

	var largest, sndlargest int
	for _, m := range monkeys {
		fmt.Printf("Monkey %d inspected items %d times \n", m.id, m.inspections)
		if m.inspections > sndlargest {
			sndlargest = m.inspections
		}
		if m.inspections > largest {
			sndlargest = largest
			largest = m.inspections
		}
	}
	fmt.Println(largest * sndlargest)

}


func (m *Monkey) throw(x *Monkey, item * Item) {
	x.items = append(x.items, item)
	// remove from monke
}

func (m *Monkey) testThrow(worryLvl *big.Int) int {
	if new(big.Int).Mod(worryLvl, big.NewInt(int64(m.test))).Cmp(big.NewInt(0)) == 0{
		//fmt.Println("trueThrow")
		return m.trueThrow
	}
	//fmt.Println("falseThrow")
	return m.falseThrow
}

func (m *Monkey) inspect(item *Item) *big.Int {
	m.inspections++
	
	if m.operationOp == "*" {
		if m.operation == 0 {
			d := new(big.Int).Mul(item.worryLvl, item.worryLvl)
			return d
		} else {
			return new(big.Int).Mul(item.worryLvl, big.NewInt(int64(m.operation)))
		}
	}

	if m.operation == 0 {
		return new(big.Int).Add(item.worryLvl, item.worryLvl)
		
	}
	return new(big.Int).Add(item.worryLvl, big.NewInt(int64(m.operation)))
}

func parseMonkey(s string) Monkey {
	var m Monkey
	lines := strings.Split(s, "\n")

	id, _ := strconv.Atoi(string(strings.Split(lines[0], " ")[1][0]))
	m.id = id

	re := regexp.MustCompile(`\d+`)
	numbers := re.FindAllString(lines[1], -1)
	for _, nr := range numbers {
		item, _ := strconv.Atoi(nr)
		m.items = append(m.items, &Item{big.NewInt(int64(item))})
	}
	operation := strings.Split(lines[2], " ")
	op := operation[6]
	m.operationOp = op
	if operation[7] == "old" {
		m.operation = 0
	} else {
		m.operation, _ = strconv.Atoi(operation[7])
	}
	divisible, _ := strconv.Atoi(re.FindString(lines[3]))
	trueThrow, _ := strconv.Atoi(re.FindString(lines[4]))
	falseThrow, _ := strconv.Atoi(re.FindString(lines[5]))
	m.test = divisible
	m.trueThrow = trueThrow
	m.falseThrow = falseThrow

	return m
}
