package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// ├───project
// │	├───file.txt (19b)
// │	└───gopher.png (70372b)
// ├───static
// │	├───a_lorem
// │	│	├───dolor.txt (empty)
// │	│	├───gopher.png (70372b)
// │	│	└───ipsum
// │	│		└───gopher.png (70372b)
// │	├───css
// │	│	└───body.css (28b)
// │	├───empty.txt (empty)
// │	├───html
// │	│	└───index.html (57b)
// │	├───js
// │	│	└───site.js (10b)
// │	└───z_lorem
// │		├───dolor.txt (empty)
// │		├───gopher.png (70372b)
// │		└───ipsum
// │			└───gopher.png (70372b)
// ├───zline
// │	├───empty.txt (empty)
// │	└───lorem
// │		├───dolor.txt (empty)
// │		├───gopher.png (70372b)
// │		└───ipsum
// │			└───gopher.png (70372b)
// └───zzfile.txt (empty)

type levelInfo struct {
	currIdx     int
	dirIdxLast  int
	fileIdxLast int
}

func recursive(parent string, level int, finLevels *map[int]levelInfo) {
	files, err := os.ReadDir(parent)
	if err != nil {
		log.Fatal(err)
	}

	newFinLevels := *finLevels
	curLevelInfo := newFinLevels[level]

	// sort filelist
	for idx, item := range files {
		if item.IsDir() {
			curLevelInfo.dirIdxLast = idx
		}
		curLevelInfo.fileIdxLast = idx
	}

	lastPrefix := "└───"

	//fmt.Println(curLevelInfo)
	for idx, item := range files {
		curLevelInfo.currIdx = idx
		newFinLevels[level] = curLevelInfo
		prefix := "├───"
		if idx == curLevelInfo.fileIdxLast {
			prefix = lastPrefix
		}

		var tabPrefix string
		for i := 0; i < level; i++ {
			if newFinLevels[i].currIdx == newFinLevels[i].fileIdxLast {
				tabPrefix += "\t"
			} else {
				tabPrefix += "│\t"
			}
		}

		if item.IsDir() {
			fmt.Printf("%s%s%s\n", tabPrefix, prefix, item.Name())
			fullPath := filepath.Join(parent, item.Name())
			recursive(fullPath, level+1, &newFinLevels)
		} else {
			fmt.Printf("%s%s%s\n", tabPrefix, prefix, item.Name())
		}

	}
}

func runme(dir string) {
	dirs := [...]string{"hello", "world", "me"}
	for _, value := range dirs {
		fmt.Println(value)
	}

}

func mapme() {
	myMap := map[int][]string{
		0: {"test", "test2"},
	}
	myMap[0] = append(myMap[0], "hello")
	fmt.Println(myMap)
}

func main() {
	finLevels := map[int]levelInfo{}
	recursive("/home/demo/go/src/hw1_tree/testdata", 0, &finLevels)
}
