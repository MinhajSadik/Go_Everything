package Basic

import (
	"fmt"
)

func randomWorker(arr []string) string {
	var worker string

	for i := 0; i < len(arr); i++ {
		worker = arr[i]
	}

	return worker
}

func Array() {
	fmt.Println("Array Learning...")

	programmers := [3]string{"minhaj", "ahmed", "sadik"}    //fixed value
	programmers1 := [...]string{"minhaj", "ahmed", "sadik"} //actual value
	// var programmers [10]string

	fmt.Println("start and end", programmers[1:2])
	fmt.Println("end", programmers[:2])
	fmt.Println("start", programmers[1:])
	fmt.Println("nothing", programmers[:])

	newWorker := randomWorker([]string{"first", "second", "third"})
	fmt.Println("newWorker:=", newWorker)

	//capacity of array
	fmt.Println("capacity", cap(programmers1))
}
