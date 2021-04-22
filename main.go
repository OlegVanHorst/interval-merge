package main

import (
	"fmt"

	"example.com/interval-merge/dataGenerator"
	"example.com/interval-merge/merge"
)

func main() {

	testData := dataGenerator.RandomDataGenerator(5, 100)
	fmt.Println(testData)

	mergedData := merge.Merge(sortedData)
	fmt.Println(mergedData)
}
