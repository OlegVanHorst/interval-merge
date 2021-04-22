package main

import (
	"fmt"

	"example.com/interval-merge/dataGenerator"
	"example.com/interval-merge/merge"
)

func main() {
	testData := dataGenerator.StaticDataGenerator()
	fmt.Println(dataGenerator.StaticDataGenerator())

	sortedData := merge.BubbleSort(testData)
	fmt.Println(sortedData)

	mergedData := merge.Merge(sortedData)
	fmt.Println(mergedData)
}
