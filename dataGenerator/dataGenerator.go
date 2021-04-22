package dataGenerator

import "math/rand"

// Creates the Example Array
func StaticDataGenerator() [][]int {
	return [][]int{{25, 30}, {2, 19}, {14, 23}, {4, 8}}
}

// Creates an Array with variable Array and Number Size and fills it with random numbers.
func RandomDataGenerator(arraySize int, numberSize int) [][]int {
	var testData [][]int
	var a int
	var b int

	for i := 0; i < arraySize; i++ {
		a = rand.Intn(numberSize)
		b = rand.Intn(numberSize)

		if a < b {
			testData = append(testData, []int{a, b})
		} else {
			testData = append(testData, []int{b, a})
		}
	}
	return testData
}
