package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Directory struct {
	name string
	parentDir *Directory
	childDir []*Directory
	files []File
	size int
}

type File struct {
	size int
	name string
}

var dir Directory
const DISCSPACE = 70000000
const NEEDED_SPACE = 30000000

func main(){
	fileP := os.Args[1]

	f, _ := os.Open(fileP)
	scanner := bufio.NewScanner(f)

	currentDir := &dir

	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '$' {
			//Parse command
			cmd := strings.Split(line, " ")[1:]
			if cmd[0] == "cd" {
				if cmd[1] == ".." {
					currentDir = currentDir.parentDir
				} else {
					newDir := &Directory{
						name: cmd[1],
						parentDir: currentDir,
					}

					currentDir.childDir = append(currentDir.childDir, newDir)
					currentDir = newDir
				}
			} else if cmd[0] == "ls" {
				
			}

		} else {
			//parse file or dir
			if line[0] != 'd' {
				file := strings.Split(line, " ")
				fileSize, _ := strconv.Atoi(file[0])
				filename := file[1]
				currentDir.files = append(currentDir.files, File{name: filename, size: fileSize})
			}
		}
	}

	sum := 0
	size := size(&dir, &sum)

	unusedSpace := DISCSPACE - size
	neededSpace := NEEDED_SPACE - unusedSpace

	smallest :=  10000000000000
	findDir(&dir, &smallest ,neededSpace)

	fmt.Println(smallest)
	
}

func findDir(dir *Directory, smallestDir *int, size int) {	
	if dir.size >= size && *smallestDir >= dir.size {
		*smallestDir = dir.size
	}

	for _, c := range dir.childDir{
		findDir(c, smallestDir, size)
	}

}

func size(dir *Directory, sum *int) int {
	fileSize := 0
	for _, f := range dir.files{
		fileSize += f.size
	}

	for _, c := range dir.childDir {
		fileSize += size(c, sum)
	}

	if fileSize <= 100000 {
		*sum += fileSize
	}
	dir.size = fileSize

	return fileSize
}

func printDir(dir Directory) {
	fmt.Println("Dir> ", dir.name, dir.size)
	for _, f := range dir.files{
		fmt.Println(f.name, f.size)
	}

	for _, c := range dir.childDir{
		printDir(*c)
	}
}