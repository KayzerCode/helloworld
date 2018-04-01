package main

import (
	"fmt"
	"log"
	"time"

	"github.com/djherbis/times"
)

var stuff = "Not ready"

func init() {
	stuff = "Ready!!!"
}

func main() {
	fmt.Printf("Hello, World, this is my first Go app\n")
	fmt.Println("We are: ", stuff)
	fmt.Println("Current time is: ", time.Now())

	processfolder := "d:/CheckAndDelete/test/"
	processfile := "resizer250318.zip"

	t, err := times.Stat(processfolder + processfile)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("File name: ", processfile)
	fmt.Println("Access Time: ", t.AccessTime())
	fmt.Println("Modification Time: ", t.ModTime())

	if t.HasChangeTime() {
		fmt.Println("Change Time: ", t.ChangeTime())
	}

	if t.HasBirthTime() {
		fmt.Println("Birth Time: ", t.BirthTime())
	}
}
