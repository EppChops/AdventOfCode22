package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var marker []byte

func main() {
	fileP := os.Args[1]

	data, _ := ioutil.ReadFile(fileP)

	index := readStream(data)

	fmt.Println(index)


}

func readStream(stream []byte) int {
	for i, b := range stream {
		if i < 14 {
			marker = append(marker, b)
			continue
		}

		if areDistinct(marker) {
			return i
		}

		marker = marker[1:]
		marker = append(marker, b)

	}
	return -1
}

func areDistinct(marker []byte) bool {
	set := make(map[byte]bool)
	
	for _, m := range marker {
		set[m] = true
	}

	return (len(set) == len(marker))
}