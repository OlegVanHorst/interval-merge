package main

import (
	"log"

	"example.com/interval-merge/dataGenerator"
	"example.com/interval-merge/merge"
)

func main() {

	testData := dataGenerator.RandomDataGenerator(5, 1000)
	log.Printf("Generated Interval: %v\n", testData)

	mergedData := merge.Merge(testData)
	log.Printf("Merged Interval: %v\n", mergedData)

}
