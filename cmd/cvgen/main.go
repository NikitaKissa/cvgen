package main

import (
	"fmt"
	"os"

	"github.com/NikitaKissa/cvgen/internal/input"
	input_json "github.com/NikitaKissa/cvgen/internal/input/json"
)

func main() {
	file, err := os.Open("./testdata/valid/cv.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var parser input.Parser = input_json.New()

	cv, err := parser.Parse(file)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", cv)
}
