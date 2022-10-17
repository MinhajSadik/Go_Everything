package main

import (
	"fmt"
	"go_everything/Algorithm"
	"go_everything/Go/Basic"
)

func main() {

	var name = "minhaj"
	const name1 = "sadik"
	profession := "programmer"
	fmt.Println(name, name1, profession)

	output := Algorithm.SequenceLength([]int{4, 7, 3, 8, 2, 1})
	fmt.Println(output)

	output = Algorithm.SequenceLength([]int{4, 7, 3, 8, 2, 1, 9, 24, 10, 11})
	fmt.Println(output)

	// Basic.Struct()
	// Basic.Switch()
	Basic.Loop()
}
