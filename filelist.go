package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var about = ""

func init() {
	about = "List all files in a folder."
}

func main() {
	fmt.Println(about)
	fmt.Println("Current time is: ", time.Now())

	var files []string
	processfolder := "d:/CheckAndDelete/"

	err := filepath.Walk(processfolder, func(path string, info os.FileInfo, err error) error {
		// No Folders in List
		if info.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	resultfile, err := os.Create("_results/files.txt")

	if err != nil {
		panic(err)
	}
	defer resultfile.Close()

	for _, file := range files {
		_, err = resultfile.WriteString(file + "\n")
		if err != nil {
			panic(err)
		}
	}
	// jsonF, err := json.Marshal(files)
	// err = ioutil.WriteFile("files.txt", jsonF, 0644)
	// if err != nil {
	// 	panic(err)
	// }

	//
	// 	fmt.Println(file)
	// }
}
