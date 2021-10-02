package main

import "fmt"

func main() {
	// // //Array in GoLang
	// var arr [4]int = [4]int{1, 2, 3, 4}
	// var arr1 = [...]int{1, 2, 3, 4, 5, 6}

	// fmt.Println("arr", arr, "arr1", arr1)

	// var arr2 = []int{1, 2, 3, 4, 5, 6}
	// fmt.Println(arr2)


	// b := "main file create for practice purpose"
	// const ami = "minhaj"
	// a := 49
	
	// fmt.Println(a, ami)
	
	// fmt.Println(b)

	// fmt.Println(utility.GetMyName())
	// fmt.Println(utility.GetName())
	// fmt.Println(controllers.Controller())
	// fmt.Println(myHelper.GetName())

	// var names = [4]string { "minhaj", "sadik", "ahmed" , "sharminhaj"}
	// var age = 24;

	// //slice (use arrays under the hood)
	// var scores = []int{1, 2, 3}
	// scores[2]  = 32
	// scores = append(scores, 98)


	// fmt.Println(scores, len(scores))

	// //slice range
	// rangeOne := names[2:3]
	// rangeTwo := names[2:]
	// rangeThree := names[:3]
	// rangeOne = append(rangeOne, "Anika")

	// fmt.Println(rangeOne, rangeTwo, rangeThree, age, names)

	var x int = 90
	var y  = &x

	fmt.Println(x)
	fmt.Println("y", *y)
}


