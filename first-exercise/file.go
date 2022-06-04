package main

import (
	"fmt"
	"os"
)

func main() {

	files := []string{"test1.txt", "test.txt", "test2.txt"}

	for index, file := range files {
		if _, err := os.Stat(file); err == nil {
			fmt.Println(index, "File", file, "exist")
		} else {
			fmt.Println(index, "File", file, "doesn't exist")
		}
	}

}
