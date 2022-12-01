package main

import (
	"fmt"
	"os"

	"github.com/bilrik/go-aoc/pkg/api"
)

func main() {
	os.Setenv("AOC_SESSION", "TestSession")
	client := api.NewClient()
	resp, err := client.PostAnswer("1", "23")
	if err != nil {
		fmt.Printf("main: could not get input data: %v\n", err)
		os.Exit(9)
	}
	fmt.Println(*resp)

	// filePath := path.Join("inputData", fmt.Sprintf("%d", client.Year), fmt.Sprintf("%d", client.Day))
	//
	// if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
	// 	err = os.MkdirAll(filePath, 0755)
	// 	if err != nil {
	// 		fmt.Printf("main: could not create directory: %v\n", err)
	// 		os.Exit(10)
	// 	}
	// } else {
	// 	fmt.Printf("main: directory already exists: %s\n", filePath)
	// }

	// if _, err := os.Stat(filePath + "/input.txt"); errors.Is(err, os.ErrNotExist) {
	// 	fmt.Printf("main: file already exists: %s\nCreating part 2 at: %s\n", filePath+"input.txt", filePath+"input2.txt")
	// 	err = os.WriteFile(filePath+"/input2.txt", []byte(*data), 0644)
	// 	if err != nil {
	// 		fmt.Printf("main: could not write input data to file: %v\n", err)
	// 		os.Exit(11)
	// 	}
	// } else if _, err := os.Stat(filePath + "input2.txt"); errors.Is(err, os.ErrNotExist) {
	// 	err = os.WriteFile(filePath+"/input.txt", []byte(*data), 0644)
	// 	if err != nil {
	// 		fmt.Printf("main: could not write input data to file: %v\n", err)
	// 		os.Exit(11)
	// 	}
	// } else {
	// 	fmt.Printf("main: you already have both files for this day: %s\n", filePath)
	// }

	os.Exit(0)
}
