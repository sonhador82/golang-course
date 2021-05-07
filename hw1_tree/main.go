package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

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
			info, err := item.Info()
			if err != nil {
				log.Fatalln(err)
			}
			var sizeInfo string
			if info.Size() > 0 {
				sizeInfo = fmt.Sprintf("%vb", info.Size())
			} else {
				sizeInfo = "empty"
			}

			fmt.Printf("%s%s%s (%s)\n", tabPrefix, prefix, item.Name(), sizeInfo)
		}

	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	finLevels := map[int]levelInfo{}
	recursive(path, 0, &finLevels)
	return nil

}

func main() {
	out := os.Stdout

	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
