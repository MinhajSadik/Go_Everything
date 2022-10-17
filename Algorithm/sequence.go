package Algorithm

import "fmt"

//Dices can have different number of sides. All sides are enumerate from 1 to n, where n is the number of sides.
//Every dice can be put with only one side up, which is its current value.
//Use as much as possible dises to make the longest possible numerical series (sequence of dice values without gaps and repetitions). Return the length of such numerical series.
//1. For set [4] the correct numerical series is {1} or {2} or {3} or {4} and the length is 1
//2. For set [4, 5, 4] the correct numerical series can be {3, 4, 5} , {2, 3, 4} the correct answer is 3
//3. For set [6, 6, 6, 6, 6, 6, 6, 6] the correct numerical series is {1, 2, 3, 4, 5, 6}, the correct answer is 6
//4. For set [4, 12, 5, 4, 3, 5, 4, 3, 8] the longest numerical series is {1 (3), 2 (3), 3 (4), 4 (4), 5 (5), 6 (8), 7 (12)} length of 7

func SequenceLength(input []int) int {
	inputMap := make(map[int]int)

	inputLen := len(input)

	for i := 0; i < inputLen; i++ {
		inputMap[input[i]] = 0
	}

	ctr := 0

	for i := 0; i < inputLen; i++ {

		currentLen := 1
		counter := 1

		for {
			value, ok := inputMap[input[i]+counter]

			if ok {
				if value != 0 {
					currentLen += value
					break
				} else {
					currentLen += 1
					counter += 1
				}
			} else {
				break
			}
		}

		fmt.Println(counter)

		if currentLen > ctr {
			ctr = currentLen
		}

		inputMap[input[i]] = currentLen
	}

	return ctr
}
