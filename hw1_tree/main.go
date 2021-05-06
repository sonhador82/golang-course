package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func recursive(parent string, dirData *map[int][]os.DirEntry, level int) {
	files, err := os.ReadDir(parent)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range files {
		if item.IsDir() {
			fullPath := filepath.Join(parent, item.Name())
			//fmt.Println(numTabs + prefix + item.Name())
			//fmt.Println(level)
			curMap := *dirData
			//fmt.Println(curMap[level])
			curMap[level] = append(curMap[level], item)
			//&dirData[level] := append(&dirData[level], item)
			level++
			recursive(fullPath, &curMap, level)
			level--
		}
	}
}

// ├───project
// ├───static
// │	├───a_lorem
// │	|   ├───a_lorem
// │	│	└───ipsum
// │	├───css
// │	├───html
// │	├───js
// │	└───z_lorem
// │		└───ipsum
// └───zline
// 	└───lorem
// 		└───ipsum

func recursive2(parent string) string {
	files, err := os.ReadDir(parent)
	if err != nil {
		log.Fatal(err)
	}

	var name string

	for _, item := range files {
		if item.IsDir() {
			fullPath := filepath.Join(parent, item.Name())
			fmt.Println(fullPath)
			name = recursive2(fullPath)

		}
	}
	return name
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	test := recursive2(path)
	fmt.Println(test)
	return nil

	dirData := map[int][]os.DirEntry{}

	recursive(path, &dirData, 0)

	for i := 0; i < len(dirData); i++ {
		//fmt.Println(i)
		var numTabs string
		for j := 0; j < i; j++ {
			numTabs += "\t"
		}
		prefix := "├───"

		numFiles := len(dirData[i]) - 1
		for kIdx, v := range dirData[i] {
			if kIdx == numFiles {
				prefix = "└───"
			}
			fmt.Fprintln(out, numTabs+prefix+v.Name())
		}
	}
	// for _, v := range dirData {
	// 	for _, item := range v {
	// 		//fmt.Println(item)
	// 	}
	// }
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
