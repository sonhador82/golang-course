package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// ├───project
// ├───static
// │	├───a_lorem
// │		├───a_lorem
// │	├───css
// │	├───html
// │	├───js

func recursive(parent string, level int) {
	files, err := os.ReadDir(parent)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range files {
		fullPath := filepath.Join(parent, item.Name())
		var numTabs string
		for i := 0; i < level; i++ {
			numTabs += "\t"
		}
		prefix := "├───"
		fmt.Println(numTabs + prefix + item.Name())
		//fmt.Println(level)
		level++
		recursive(fullPath, level)
		level--
	}

	//fmt.Println(files)

	// if idx == 0 {
	// 	return
	// }
	// idx--
	//fmt.Println(idx)
	//recursive(idx)
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
	fmt.Println("JHe")
	//recursive("/home/demo/go/src/hw1_tree/testdata", 0)
	mapme()
}
