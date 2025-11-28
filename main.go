package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("course.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var courses map[string][]any
	if err := json.NewDecoder(file).Decode(&courses); err != nil {
		panic(err)
	}
	for k, v := range courses {
		fmt.Printf("%q - %q\n", k, v)
	}
}
