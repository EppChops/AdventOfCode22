package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fileP := os.Args[1]

	f, _ := os.Open(fileP)
	scanner := bufio.NewScanner(f)
	
	trees := make([][]int, 0)
	for scanner.Scan() {
		treeRow := make([]int, 0)

		for _, b := range scanner.Text(){
			digit, _ := strconv.Atoi(string(b))
			treeRow = append(treeRow, digit)
		}

		trees = append(trees, treeRow)
		

	}

//	fmt.Println(trees)

	count := 0

	topScore := 0
	for i := 1; i < len(trees) - 1; i++{
		for j := 1; j < len(trees) -1; j++ {
			if checkLeftIsVisible(trees, i, j) || 
			checkRightIsVisible(trees, i, j) ||
			checkUpIsVisible(trees, i, j) ||
			checkDownIsVisible(trees, i, j){
				count++


				score := checkLeftScore(trees, i, j)
				score *= checkRightScore(trees, i, j)
				score *= checkUpScore(trees, i, j)
				score *= checkDownScore(trees, i, j)
				fmt.Println(score)
				if score > topScore {
					topScore = score
				}
			}
		}
	}

	count += len(trees) * 2
	count += (len(trees)-2) * 2  
	fmt.Println(count)
	fmt.Println("score:")
	fmt.Println(topScore)
}

func checkLeftIsVisible(trees [][]int, row int, col int) bool {
	treeHeight := trees[row][col]

	for i := 0; i < col; i++{
		if trees[row][i] >= treeHeight {
			
			return false
		}
	}
	return true
}

func checkRightIsVisible(trees [][]int, row int, col int) bool {
	treeHeight := trees[row][col]

	for i := len(trees) -1; i > col; i-- {
		if trees[row][i] >= treeHeight {
			
			return false
		}
	}
	return true
}

func checkUpIsVisible(trees [][]int, row int, col int) bool {
	treeHeight := trees[row][col]

	for i := 0; i < row; i++ {
		if trees[i][col] >= treeHeight {
			
			return false
		}
	}
	return true
}

func checkDownIsVisible(trees [][]int, row int, col int) bool {
	treeHeight := trees[row][col]

	for i := len(trees) - 1; i >= row; i-- {
		if trees[i][col] >= treeHeight {
			
			return false
		}
	}
	return true
}

func checkLeftScore(trees [][]int, row int, col int) int {
	treeHeight := trees[row][col]

	score := 0
	for i := col -1; i >= 0; i--{
		score++
		if trees[row][i] >= treeHeight {
			return score
		}
		
	}

	return score
}

func checkRightScore(trees [][]int, row int, col int) int {
	treeHeight := trees[row][col]
	score := 0
	for i := col+1; i <= len(trees) -1; i++{
		score++
		if trees[row][i] >= treeHeight {
			return score
		}
		
	}
	return score
}

func checkUpScore(trees [][]int, row int, col int) int {
	treeHeight := trees[row][col]

	score := 0
	for i := row-1; i >= 0; i-- {
		score++
		if trees[i][col] >= treeHeight {
			
			return score
		}
	}
	return score
}

func checkDownScore(trees [][]int, row int, col int) int {
	treeHeight := trees[row][col]
	score := 0
	for i := row+1; i <= len(trees)-1; i++ {
		score++
		if trees[i][col] >= treeHeight {
			
			return score
		}
	}
	return score
}