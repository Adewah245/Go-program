package main

import (
	"fmt"
	"os"
)

func reader(input string) string {
	data, _:= os.ReadFile("input.txt")
	input = string(data)
	fmt.Println(input)

	result := processor(inputs(reader(input)))
		os.WriteFile("history.txt", []byte(result), 0644)
		fmt.Println(result)
	return input
}